-- +goose Up
-- +goose StatementBegin
CREATE materialized view IF NOT EXISTS dsp_wins_view TO dsp_wins AS
SELECT
    JSONExtractString(msg, 'requestId') as requestId,
    JSONExtractString(msg, 'dsp') as dsp,
    toDecimal32(JSONExtractString(msg, 'price'), 4) as price
FROM dsp_wins_consumer;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS dsp_wins_view
-- +goose StatementEnd
