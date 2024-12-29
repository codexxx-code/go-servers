-- +goose Up
-- +goose StatementBegin
ALTER TABLE ssp.ssps ADD is_enable bool DEFAULT true NOT NULL;
COMMENT ON COLUMN ssp.ssps.is_enable IS 'Индикатор, будет ли SSP принимать трафик';
ALTER TABLE ssp.ssps ALTER COLUMN is_enable DROP DEFAULT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE ssp.ssps DROP COLUMN is_enable;
-- +goose StatementEnd
