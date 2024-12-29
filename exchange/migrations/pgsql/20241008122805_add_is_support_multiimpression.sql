-- +goose Up
-- +goose StatementBegin
ALTER TABLE ssp.dsps ADD is_support_multiimpression bool DEFAULT false NOT NULL;
COMMENT ON COLUMN ssp.dsps.is_support_multiimpression IS 'Умеет ли принимать DSP несколько импрешенов в одном запросе';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE ssp.dsps DROP COLUMN is_support_multiimpression;
-- +goose StatementEnd
