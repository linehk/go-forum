package model

import (
	"time"
)

type Thread struct {
	ID        int
	CreatedAt time.Time
	UUID      string
	Subject   string
	UserID    int
}
