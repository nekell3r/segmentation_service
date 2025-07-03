package domain

// UserRepository определяет методы для работы с пользователями

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id int64) (*User, error)
	Create(user *User) error
}

// SegmentRepository определяет методы для работы с сегментами

type SegmentRepository interface {
	GetAllSegments() ([]Segment, error)
	GetSegmentByName(name string) (*Segment, error)
	CreateSegment(segment *Segment) error
	DeleteSegment(name string) error
	RenameSegment(oldName, newName string) error
	AddUserToSegment(userID int64, segmentName string) error
	RemoveUserFromSegment(userID int64, segmentName string) error
	GetUserSegments(userID int64) ([]Segment, error)
	DistributeSegmentToPercent(segmentName string, percent float64) error
}

// SegmentService определяет бизнес-логику для сегментов

type SegmentService interface {
	CreateSegment(name string) error
	DeleteSegment(name string) error
	RenameSegment(oldName, newName string) error
	AddUserToSegment(userID int64, segmentName string) error
	RemoveUserFromSegment(userID int64, segmentName string) error
	DistributeSegmentToPercent(segmentName string, percent float64) error
	GetUserSegments(userID int64) ([]Segment, error)
}
