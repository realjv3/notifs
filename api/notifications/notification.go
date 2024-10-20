package notifications

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID          string
	UserId      uuid.UUID `db:"userId"`
	Title       string
	Description string
	Type        string
	CreatedAt   time.Time `db:"createdAt"`
}
