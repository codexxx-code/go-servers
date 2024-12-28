-- +goose Up
-- +goose StatementBegin
ALTER SCHEMA zodiac RENAME TO horoscope;
ALTER TABLE horoscope.forecasts RENAME TO horoscopes;
CREATE SCHEMA generator;
ALTER TABLE horoscope.prompts SET SCHEMA generator;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS horoscope CASCADE;
DROP SCHEMA IF EXISTS generator CASCADE;
-- +goose StatementEnd
