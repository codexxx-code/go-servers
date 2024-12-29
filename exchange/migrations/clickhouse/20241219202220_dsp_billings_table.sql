-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS dsp_billings
(
    exchange_bid_id String,
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
    ssp_request_bid_floor_in_default_currency DECIMAL(10, 6),
    dsp_request_id String,
    dsp_request_impression_id String,
    dsp_request_slug String,
    dsp_request_domain LowCardinality(String),
    dsp_request_bundle LowCardinality(String),
    dsp_request_geo_country LowCardinality(String),
    dsp_request_geo_region LowCardinality(String),
    dsp_request_geo_city LowCardinality(String),
    dsp_request_width Int,
    dsp_request_height Int,
    dsp_request_ad_type LowCardinality(String),
    dsp_request_bid_floor_in_default_currency DECIMAL(10, 6),
    dsp_response_price_in_default_currency DECIMAL(10, 6),
    ssp_response_price_in_default_currency DECIMAL(10, 6),
    fact_price_ssp_in_default_currency DECIMAL(10, 6),
    fact_price_dsp_in_default_currency DECIMAL(10, 6)
) ENGINE = MergeTree()
      PARTITION BY toYYYYMMDD(request_date_time)
      ORDER BY (request_date_time)
      TTL request_date_time + INTERVAL 1 DAY;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dsp_billings;
-- +goose StatementEnd
