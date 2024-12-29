package service

import (
	"context"

	"golang.org/x/sync/singleflight"

	"exchange/internal/services/fraudScore/model"
)

// FraudScoreService представляет бизнес-логику для работы с Fraudscore API.
type FraudScoreService struct {
	client       FraudscoreNetwork
	singleflight singleflight.Group
}

type FraudscoreNetwork interface {
	CheckFraudScore(_ context.Context, req model.IsFraudReq) (bool, error)
}

// NewFraudscoreService создает новый сервис FraudScoreService.
func NewFraudscoreService(client FraudscoreNetwork) *FraudScoreService {
	return &FraudScoreService{
		client: client,
	}
}
