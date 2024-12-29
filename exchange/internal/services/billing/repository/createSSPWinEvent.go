package repository

import (
	"context"
	"time"

	"exchange/internal/config"
)

type CreateSSPEventWinReq struct {
	RequestID string    `json:"requestId"`
	SSP       string    `json:"ssp"`
	Price     string    `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

func (r *BillingRepository) CreateSSPWinEvent(_ context.Context, win CreateSSPEventWinReq) error {
	win.Timestamp = time.Now()
	return r.writeToTopic(win, config.Load().Queue.Topic.OldAnalytic.SSPWins)
}
