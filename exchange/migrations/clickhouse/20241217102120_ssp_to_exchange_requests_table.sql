-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ssp_to_exchange_requests
(
    exchange_impression_id String,
    exchange_id String,
    request_date_time DateTime,
    ssp_request_id String,
    ssp_request_impression_id String,
    ssp_request_slug String,
    ssp_request_domain LowCardinality(String),
    ssp_request_bundle LowCardinality(String),
    ssp_request_geo_country LowCardinality(String),
    ssp_request_geo_region LowCardinality(String),
    ssp_request_geo_city LowCardinality(String),
    ssp_request_width Int,
    ssp_request_height Int,
    ssp_request_ad_type LowCardinality(String),
    ssp_request_bid_floor_in_default_currency DECIMAL(10, 6)
) ENGINE = MergeTree()
      PARTITION BY toYYYYMMDD(request_date_time)
      ORDER BY (request_date_time)
      TTL request_date_time + INTERVAL 1 DAY;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ssp_to_exchange_requests;
-- +goose StatementEnd
