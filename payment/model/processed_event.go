package model

import (
	"time"
)

type ProcessedEvent struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
}
