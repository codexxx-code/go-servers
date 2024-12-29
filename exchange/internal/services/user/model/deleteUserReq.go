package model

import (
	"time"
)

type DeleteUserReq struct {
	ID        string
	DeletedAt time.Time
	DeletedBy string
}
