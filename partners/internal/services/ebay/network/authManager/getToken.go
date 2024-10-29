package authManager

import (
	"context"
	"time"

	"partners/internal/services/ebay/network/model"
)

const bearerTokenTTL = 7000 * time.Second

func (s *EbayAuthManager) GetToken(ctx context.Context) (string, error) {

	// Блокируем токен
	s.bearer.mu.Lock()
	defer s.bearer.mu.Unlock()

	// Если токен не истек
	if !time.Now().After(s.bearer.bearerTokenExpiredDate) {

		// Просто возвращаем токен
		return s.bearer.bearerToken, nil
	}

	// Запрашиваем новый токен
	authRes, err := s.ebayNetwork.Auth(ctx, model.AuthReq{
		BasicToken: s.basicToken,
	})
	if err != nil {
		return "", err
	}

	// Обновляем токен
	s.bearer.bearerToken = authRes.AccessToken
	s.bearer.bearerTokenExpiredDate = time.Now().Add(bearerTokenTTL)

	return s.bearer.bearerToken, nil
}
