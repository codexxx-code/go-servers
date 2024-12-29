package model

import (
	"time"

	"git.yabbi.me/ssp/api-ssp-server-go/proto"

	"pkg/errors"
)

type DeleteUserReq struct {
	ID        string
	DeletedAt time.Time
	DeletedBy string
}

type ProtoDeleteUserReq struct {
	*proto.DeleteUserRequest
}

func (r ProtoDeleteUserReq) ConvertToModel() (res DeleteUserReq, err error) {

	// Проверяем, что запрос не nil
	if r.DeleteUserRequest == nil {
		return res, errors.BadRequest.New("DeleteUserRequest is required")
	}

	// Маппим структуру
	return DeleteUserReq{
		ID:        r.Uuid,
		DeletedAt: time.Time{}, // Заполняется после маппинга
		DeletedBy: "",          // Заполняется после маппинга
	}, nil
}
