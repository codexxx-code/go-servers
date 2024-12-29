-- +goose Up
-- +goose StatementBegin

-- Создание таблицы source_traffic_types, если она ещё не создана
CREATE TABLE IF NOT EXISTS ssp.source_traffic_types
(
    slug VARCHAR NOT NULL,
    CONSTRAINT source_traffic_types_pk PRIMARY KEY (slug)
);

-- Добавление данных в таблицу source_traffic_types, если они ещё не добавлены
INSERT INTO ssp.source_traffic_types (slug)
VALUES ('in_app'),
       ('mobile_web'),
       ('desktop'),
       ('smart')
ON CONFLICT DO NOTHING;

-- Создание таблицы integration_types, если она ещё не создана
CREATE TABLE IF NOT EXISTS ssp.integration_types
(
    slug VARCHAR NOT NULL,
    CONSTRAINT integration_types_pk PRIMARY KEY (slug)
);

-- Добавление данных в таблицу integration_types, если они ещё не добавлены
INSERT INTO ssp.integration_types (slug)
VALUES ('ortb'),
       ('vast'),
       ('feed')
ON CONFLICT DO NOTHING;

-- Добавление столбцов в таблицу ssp.dsps
ALTER TABLE ssp.dsps
    ADD is_enable           BOOLEAN NOT NULL DEFAULT FALSE,
    ADD integration_type    VARCHAR NOT NULL DEFAULT 'ortb',
    ADD source_traffic_type VARCHAR NOT NULL DEFAULT 'in_app';

-- Связывание таблицы dsp и source_traffic_types
ALTER TABLE ssp.dsps
    ADD CONSTRAINT dsps_source_traffic_types_fk FOREIGN KEY (source_traffic_type) REFERENCES ssp.source_traffic_types (slug);

-- Связывание таблицы dsp и integration_types
ALTER TABLE ssp.dsps
    ADD CONSTRAINT dsps_integration_types_fk FOREIGN KEY (integration_type) REFERENCES ssp.integration_types (slug);

-- Создание таблицы format_types, если она ещё не создана
CREATE TABLE IF NOT EXISTS ssp.format_types
(
    slug VARCHAR NOT NULL,
    CONSTRAINT format_types_pk PRIMARY KEY (slug)
);

-- Добавление данных в таблицу format_types, если они ещё не добавлены
INSERT INTO ssp.format_types (slug)
VALUES ('banner'),
       ('video'),
       ('push'),
       ('clickunder'),
       ('native')
ON CONFLICT DO NOTHING;

-- Создание связующей таблицы dsp и format_types
CREATE TABLE ssp.dsps_to_format_types
(
    format_type_slug VARCHAR NOT NULL,
    dsp_slug         VARCHAR NOT NULL,
    CONSTRAINT dsps_to_format_types_pk PRIMARY KEY (format_type_slug, dsp_slug),
    CONSTRAINT dsps_to_format_types_format_types_fk FOREIGN KEY (format_type_slug) REFERENCES ssp.format_types (slug),
    CONSTRAINT dsps_to_format_types_dsps_fk FOREIGN KEY (dsp_slug) REFERENCES ssp.dsps (slug) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ssp.dsps_to_format_types;
DROP TABLE ssp.format_types;
DROP TABLE ssp.source_traffic_types;
DROP TABLE ssp.integration_types;
DROP TABLE ssp.dsps;
-- +goose StatementEnd
