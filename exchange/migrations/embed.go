package migrations

import "embed"

//go:embed pgsql/*.sql
var EmbedMigrationsPgsql embed.FS

//go:embed clickhouse/*.sql
var EmbedMigrationsClickhouse embed.FS
