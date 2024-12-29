-- +goose Up
-- +goose StatementBegin
CREATE materialized view IF NOT EXISTS dsp_to_exchange_responses_view TO dsp_to_exchange_responses AS
SELECT
    JSONExtractString(msg, 'exchange_bid_id') as exchange_bid_id,
    JSONExtractString(msg, 'exchange_impression_id') as exchange_impression_id,
    JSONExtractString(msg, 'exchange_id') as exchange_id,
    toDateTime(JSONExtractString(msg, 'request_date_time')) as request_date_time,
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
    JSONExtractString(msg, 'dsp_request_impression', 'request_id') as dsp_request_id,
    JSONExtractString(msg, 'dsp_request_impression', 'impression_id') as dsp_request_impression_id,
    JSONExtractString(msg, 'dsp_request_impression', 'slug') as dsp_request_slug,
    JSONExtractString(msg, 'dsp_request_impression', 'domain') as dsp_request_domain,
    JSONExtractString(msg, 'dsp_request_impression', 'bundle') as dsp_request_bundle,
    JSONExtractString(msg, 'dsp_request_impression', 'geo', 'country') as dsp_request_geo_country,
    JSONExtractString(msg, 'dsp_request_impression', 'geo', 'region') as dsp_request_geo_region,
    JSONExtractString(msg, 'dsp_request_impression', 'geo', 'city') as dsp_request_geo_city,
    JSONExtractInt(msg, 'dsp_request_impression', 'width') as dsp_request_width,
    JSONExtractInt(msg, 'dsp_request_impression', 'height') as dsp_request_height,
    JSONExtractString(msg, 'dsp_request_impression', 'ad_type') as dsp_request_ad_type,
    toDecimal32(JSONExtractString(msg, 'dsp_request_impression', 'bid_floor_in_default_currency'), 6) as dsp_request_bid_floor_in_default_currency,
    toDecimal32(JSONExtractString(msg, 'dsp_response_bid', 'price_in_default_currency'), 6) as dsp_response_price_in_default_currency
FROM dsp_to_exchange_responses_consumer;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS dsp_to_exchange_responses_view
-- +goose StatementEnd
