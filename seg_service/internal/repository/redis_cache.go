package repository

import (
	"context"
	"seg_service/internal/domain"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	rdb *redis.Client
}

func NewRedisCache(rdb *redis.Client) *RedisCache {
	return &RedisCache{rdb: rdb}
}

// Методы для кеширования сегментов пользователя будут добавлены позже

// InvalidateUserSegments — удаляет кеш сегментов пользователя
func (c *RedisCache) InvalidateUserSegments(ctx context.Context, userID int64) error {
	return c.rdb.Del(ctx, c.userSegmentsKey(userID)).Err()
}

// GetUserSegments — получает сегменты пользователя из кеша
func (c *RedisCache) GetUserSegments(ctx context.Context, userID int64) ([]domain.Segment, bool) {
	// Заглушка: всегда cache miss
	return nil, false
}

// SetUserSegments — сохраняет сегменты пользователя в кеш
func (c *RedisCache) SetUserSegments(ctx context.Context, userID int64, segments []domain.Segment) error {
	// Заглушка: не сохраняет
	return nil
}

func (c *RedisCache) userSegmentsKey(userID int64) string {
	return "user_segments:" + string(rune(userID))
}
