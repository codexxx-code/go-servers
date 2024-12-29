-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ssp_to_exchange_requests_by_hour
(
    request_hour DateTime,
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
    ssp_requests_count Int
) ENGINE = SummingMergeTree()
      PARTITION BY toYYYYMMDD(request_hour)
      ORDER BY (
        request_hour,
        ssp_request_slug,
        ssp_request_domain,
        ssp_request_bundle,
        ssp_request_geo_country,
        ssp_request_geo_region,
        ssp_request_geo_city,
        ssp_request_width,
        ssp_request_height,
        ssp_request_ad_type
      );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ssp_to_exchange_requests_by_hour;
-- +goose StatementEnd
