-- +goose Up
-- +goose StatementBegin

ALTER TABLE ssp.ssps ALTER COLUMN clickunder_drum_size DROP DEFAULT;
ALTER TABLE ssp.ssps ALTER COLUMN integration_type DROP DEFAULT;
ALTER TABLE ssp.ssps ALTER COLUMN endpoint_url DROP DEFAULT;
ALTER TABLE ssp.ssps ALTER COLUMN source_traffic_type DROP DEFAULT;
ALTER TABLE ssp.ssps ALTER COLUMN billing_type DROP DEFAULT;
ALTER TABLE ssp.ssps ALTER COLUMN auction_second_price DROP DEFAULT;
ALTER TABLE ssp.ssps ALTER COLUMN currency DROP DEFAULT;
ALTER TABLE ssp.dsps ALTER COLUMN is_support_multiimpression DROP DEFAULT;
ALTER TABLE ssp.users ALTER COLUMN is_deleted DROP DEFAULT;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
