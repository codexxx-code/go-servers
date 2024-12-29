-- +goose Up
-- +goose StatementBegin
ALTER TABLE ssp.ssps ADD multiplication_factor int4 DEFAULT 1 NOT NULL;
COMMENT ON COLUMN ssp.ssps.multiplication_factor IS 'Количество исходящих ответов к DSP на один входящий запрос от SSP';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE ssp.ssps DROP COLUMN multiplication_factor;
-- +goose StatementEnd
