package model

import (
	"time"
)

type UpdateLastLoginAtReq struct {
	ID          string
	LastLoginAt time.Time
}
