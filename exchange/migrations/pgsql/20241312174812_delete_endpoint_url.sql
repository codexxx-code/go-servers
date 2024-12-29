-- +goose Up
-- +goose StatementBegin
ALTER TABLE ssp.ssps
DROP COLUMN endpoint_url;

-- +goose StatementEnd

