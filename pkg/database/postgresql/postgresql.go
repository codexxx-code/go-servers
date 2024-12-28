package postgresql

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" //nolint:golint

	"pkg/sql"
)

type PostgreSQLConfig struct {
	Host     string `env:"PGSQL_HOST,required"`
	User     string `env:"PGSQL_USER,required"`
	Password string `env:"PGSQL_PASSWORD,required"`
}

func (c *PostgreSQLConfig) GetURL(databaseName string) string {
	return fmt.Sprintf("postgres://%v:%v@%v/%v", c.User, c.Password, c.Host, databaseName)
}

func NewClientSQL(repo PostgreSQLConfig, databaseName string) (*sql.DB, error) {
	db, err := sql.Open("pgx", repo.GetURL(databaseName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db.Unsafe(), nil
}
