package service

import (
	"context"
	"math/rand"
	"seg_service/internal/domain"
	"seg_service/internal/repository"
)

type SegmentServiceImpl struct {
	users    domain.UserRepository
	segments domain.SegmentRepository
	cache    *repository.RedisCache
}

func NewSegmentService(users domain.UserRepository, segments domain.SegmentRepository, cache *repository.RedisCache) *SegmentServiceImpl {
	return &SegmentServiceImpl{
		users:    users,
		segments: segments,
		cache:    cache,
	}
}

func (s *SegmentServiceImpl) CreateSegment(name string) error {
	return s.segments.CreateSegment(&domain.Segment{Name: name})
}

func (s *SegmentServiceImpl) DeleteSegment(name string) error {
	return s.segments.DeleteSegment(name)
}

func (s *SegmentServiceImpl) RenameSegment(oldName, newName string) error {
	return s.segments.RenameSegment(oldName, newName)
}

func (s *SegmentServiceImpl) AddUserToSegment(userID int64, segmentName string) error {
	err := s.segments.AddUserToSegment(userID, segmentName)
	if err == nil && s.cache != nil {
		_ = s.cache.InvalidateUserSegments(context.Background(), userID)
	}
	return err
}

func (s *SegmentServiceImpl) RemoveUserFromSegment(userID int64, segmentName string) error {
	err := s.segments.RemoveUserFromSegment(userID, segmentName)
	if err == nil && s.cache != nil {
		_ = s.cache.InvalidateUserSegments(context.Background(), userID)
	}
	return err
}

func (s *SegmentServiceImpl) DistributeSegmentToPercent(segmentName string, percent float64) error {
	users, err := s.users.GetAll()
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return nil
	}
	count := int(float64(len(users)) * percent / 100.0)
	if count == 0 {
		return nil
	}
	ids := make([]int64, len(users))
	for i, u := range users {
		ids[i] = u.ID
	}
	rand.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })
	for i := 0; i < count; i++ {
		_ = s.AddUserToSegment(ids[i], segmentName)
	}
	return nil
}

func (s *SegmentServiceImpl) GetUserSegments(userID int64) ([]domain.Segment, error) {
	if s.cache != nil {
		if segments, ok := s.cache.GetUserSegments(context.Background(), userID); ok {
			return segments, nil
		}
	}
	segments, err := s.segments.GetUserSegments(userID)
	if err == nil && s.cache != nil {
		_ = s.cache.SetUserSegments(context.Background(), userID, segments)
	}
	return segments, err
}
