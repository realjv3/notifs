package notifications

import (
	"github.com/realjv3/homework/users"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {

	return &Service{db}
}

// GetNotificationsByPrefType fetches the passed user's notifications by their preferred type
// TODO accommodate pagination
func (s *Service) GetNotificationsByPrefType(user *users.User) ([]Notification, error) {
	var notifications []Notification

	err := s.db.Select(
		&notifications,
		"SELECT * FROM notifications WHERE userId = ? AND type = ?",
		user.ID.String(),
		user.Preference,
	)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}
