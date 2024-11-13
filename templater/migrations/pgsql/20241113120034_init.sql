-- +goose Up
-- +goose StatementBegin

-- Создаем схему для хранения таблиц
CREATE SCHEMA templater;

-- Создаем таблицу для хранения шаблонов
CREATE TABLE templater.templates (
	id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
	ssp_slug varchar NOT NULL,
	"template" varchar NOT NULL,
	CONSTRAINT templates_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA templater;
-- +goose StatementEnd
