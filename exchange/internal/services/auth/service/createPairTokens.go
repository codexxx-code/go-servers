package service

import (
	"exchange/internal/enum/permission"
	authModel "exchange/internal/services/auth/model"
	"pkg/jwtManager"
)

// TODO: Обратить внимание, что по access токену можно будет обновить пару токенов, что сводит на нет использование refresh токена
// Обратная ситуация: по refresh токену можно обращаться к api и запрос будет проходить
// createPairTokens создает пару токенов с пермишенами и идентификатором
func createPairTokens(permissions []permission.Permission, id string) (tokens authModel.Tokens, err error) {

	claims := authModel.Claims{
		Permissions: permissions,
		ID:          id,
	}

	tokens.AccessToken, err = jwtManager.GenerateToken(jwtManager.AccessToken, claims)
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken, err = jwtManager.GenerateToken(jwtManager.RefreshToken, claims)
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}
