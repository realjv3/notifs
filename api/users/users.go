package users

import (
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {

	return &Service{db}
}

func (s *Service) GetUserByEmail(email string) (*User, error) {
	var user User

	err := s.db.Get(&user, "SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) SetUserNotificationPref(user *User) error {
	_, err := s.db.Exec(`UPDATE users SET preference = ? WHERE id = ?`, user.Preference, user.ID.String())
	if err != nil {
		return err
	}

	return nil
}
