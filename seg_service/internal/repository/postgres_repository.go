package repository

import (
	"database/sql"
	"fmt"

	"seg_service/internal/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// --- UserRepository ---
func (r *PostgresRepository) GetAll() ([]domain.User, error) {
	rows, err := r.db.Query("SELECT id FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *PostgresRepository) GetByID(id int64) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id FROM users WHERE id = $1", id)
	var u domain.User
	if err := row.Scan(&u.ID); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *PostgresRepository) Create(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (id) VALUES ($1) ON CONFLICT DO NOTHING", user.ID)
	return err
}

// --- SegmentRepository ---
func (r *PostgresRepository) GetAllSegments() ([]domain.Segment, error) {
	rows, err := r.db.Query("SELECT name FROM segments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var segments []domain.Segment
	for rows.Next() {
		var s domain.Segment
		if err := rows.Scan(&s.Name); err != nil {
			return nil, err
		}
		segments = append(segments, s)
	}
	return segments, nil
}

func (r *PostgresRepository) GetSegmentByName(name string) (*domain.Segment, error) {
	row := r.db.QueryRow("SELECT name FROM segments WHERE name = $1", name)
	var s domain.Segment
	if err := row.Scan(&s.Name); err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *PostgresRepository) CreateSegment(segment *domain.Segment) error {
	_, err := r.db.Exec("INSERT INTO segments (name) VALUES ($1) ON CONFLICT DO NOTHING", segment.Name)
	return err
}

func (r *PostgresRepository) DeleteSegment(name string) error {
	_, err := r.db.Exec("DELETE FROM segments WHERE name = $1", name)
	return err
}

func (r *PostgresRepository) RenameSegment(oldName, newName string) error {
	_, err := r.db.Exec("UPDATE segments SET name = $1 WHERE name = $2", newName, oldName)
	return err
}

func (r *PostgresRepository) AddUserToSegment(userID int64, segmentName string) error {
	_, err := r.db.Exec("INSERT INTO user_segments (user_id, segment_name) VALUES ($1, $2) ON CONFLICT DO NOTHING", userID, segmentName)
	return err
}

func (r *PostgresRepository) RemoveUserFromSegment(userID int64, segmentName string) error {
	_, err := r.db.Exec("DELETE FROM user_segments WHERE user_id = $1 AND segment_name = $2", userID, segmentName)
	return err
}

func (r *PostgresRepository) GetUserSegments(userID int64) ([]domain.Segment, error) {
	rows, err := r.db.Query("SELECT segment_name FROM user_segments WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var segments []domain.Segment
	for rows.Next() {
		var s domain.Segment
		if err := rows.Scan(&s.Name); err != nil {
			return nil, err
		}
		segments = append(segments, s)
	}
	return segments, nil
}

func (r *PostgresRepository) DistributeSegmentToPercent(segmentName string, percent float64) error {
	// Заглушка: реализация будет в сервисе
	return fmt.Errorf("not implemented")
}
