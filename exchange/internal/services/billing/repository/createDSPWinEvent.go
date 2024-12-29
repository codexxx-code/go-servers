package repository

import (
	"context"
	"time"

	"exchange/internal/config"
)

type CreateDSPEventWinReq struct {
	RequestID string    `json:"requestId"`
	DSP       string    `json:"dsp"`
	Price     string    `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

func (r *BillingRepository) CreateDSPWinEvent(_ context.Context, win CreateDSPEventWinReq) error {
	win.Timestamp = time.Now()
	return r.writeToTopic(win, config.Load().Queue.Topic.OldAnalytic.SSPWins)
}
