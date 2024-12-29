-- +goose Up
-- +goose StatementBegin
CREATE materialized view IF NOT EXISTS ssp_wins_view TO ssp_wins AS
SELECT
    JSONExtractString(msg, 'requestId') as requestId,
    JSONExtractString(msg, 'ssp') as ssp,
    toDecimal32(JSONExtractString(msg, 'price'), 4) as price
FROM ssp_wins_consumer;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS ssp_wins_view
-- +goose StatementEnd
