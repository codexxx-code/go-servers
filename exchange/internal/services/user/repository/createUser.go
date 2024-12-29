package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"

	userPermission "exchange/internal/enum/permission"
	"exchange/internal/services/user/repository/userDDL"
	"pkg/uuid"
)

type CreateUserReq struct {
	ID           string
	LastName     string
	FirstName    string
	Email        string
	PasswordHash []byte
	PasswordSalt []byte
	Permissions  []userPermission.Permission
	AuthorID     *string
}

// CreateUser Создает нового пользователя
func (r *UserRepository) CreateUser(ctx context.Context, user CreateUserReq) (string, error) {

	// TODO: Добавить транзакцию

	user.ID = uuid.New()

	// Создаем пользователя в базе данных
	return user.ID, r.pgsql.Exec(ctx, squirrel.
		Insert(userDDL.Table).
		SetMap(map[string]any{
			userDDL.ColumnID:           user.ID,
			userDDL.ColumnLastName:     user.LastName,
			userDDL.ColumnFirstName:    user.FirstName,
			userDDL.ColumnEmail:        user.Email,
			userDDL.ColumnPasswordHash: user.PasswordHash,
			userDDL.ColumnPasswordSalt: user.PasswordSalt,
			userDDL.ColumnAuthorID:     user.AuthorID,
			userDDL.ColumnLastLoginAt:  nil,
			userDDL.ColumnIsDeleted:    false,
			userDDL.ColumnPermissions:  pq.Array(user.Permissions),
		}),
	)
}
