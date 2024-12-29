-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS dsp_wins
(
    requestId  String,
    dsp LowCardinality(String), -- https://clickhouse.com/docs/ru/sql-reference/data-types/lowcardinality
    price String,
    created_at DateTime DEFAULT now()
    ) ENGINE = ReplacingMergeTree() -- https://clickhouse.com/docs/ru/engines/table-engines/mergetree-family/replacingmergetree
      PARTITION BY toYYYYMMDD(created_at)
      ORDER BY (requestId);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dsp_wins;
-- +goose StatementEnd
