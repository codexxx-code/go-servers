-- +goose Up
-- +goose StatementBegin
CREATE materialized view IF NOT EXISTS ssp_to_exchange_requests_by_hour_view TO ssp_to_exchange_requests_by_hour AS
SELECT
    JSONExtractString(msg, 'exchange_impression_id') as exchange_impression_id,
    JSONExtractString(msg, 'exchange_id') as exchange_id,
    toStartOfHour(toDateTime(JSONExtractString(msg, 'request_date_time'))) as request_hour,
    JSONExtractString(msg, 'ssp_request_impression', 'request_id') as ssp_request_id,
    JSONExtractString(msg, 'ssp_request_impression', 'impression_id') as ssp_request_impression_id,
    JSONExtractString(msg, 'ssp_request_impression', 'slug') as ssp_request_slug,
    JSONExtractString(msg, 'ssp_request_impression', 'domain') as ssp_request_domain,
    JSONExtractString(msg, 'ssp_request_impression', 'bundle') as ssp_request_bundle,
    JSONExtractString(msg, 'ssp_request_impression', 'geo', 'country') as ssp_request_geo_country,
    JSONExtractString(msg, 'ssp_request_impression', 'geo', 'region') as ssp_request_geo_region,
    JSONExtractString(msg, 'ssp_request_impression', 'geo', 'city') as ssp_request_geo_city,
    JSONExtractInt(msg, 'ssp_request_impression', 'width') as ssp_request_width,
    JSONExtractInt(msg, 'ssp_request_impression', 'height') as ssp_request_height,
    JSONExtractString(msg, 'ssp_request_impression', 'ad_type') as ssp_request_ad_type,
    toDecimal32(JSONExtractString(msg, 'ssp_request_impression', 'bid_floor_in_default_currency'), 6) as ssp_request_bid_floor_in_default_currency,
    1 as ssp_requests_count
FROM ssp_to_exchange_requests_consumer;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS ssp_to_exchange_requests_by_hour_view;
-- +goose StatementEnd
