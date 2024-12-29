package userDDL

import "exchange/internal/ddl"

const (
	Table          = ddl.SchemaSSP + "." + "users"
	TableWithAlias = Table + " " + alias
	alias          = "u"
)

const (
	ColumnID           = "id"
	ColumnLastName     = "last_name"
	ColumnFirstName    = "first_name"
	ColumnEmail        = "email"
	ColumnPasswordHash = "password_hash"
	ColumnPasswordSalt = "password_salt"
	ColumnAuthorID     = "author_id"
	ColumnLastLoginAt  = "last_login_at"
	ColumnIsDeleted    = "is_deleted"
	ColumnPermissions  = "permissions"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
