-- +goose Up
-- +goose StatementBegin

-- Добавление столбцов в таблицу ssp.ssps
ALTER TABLE ssp.ssps
    ADD integration_type VARCHAR NOT NULL DEFAULT 'ortb',
    ADD endpoint_url VARCHAR NOT NULL DEFAULT '',
    ADD source_traffic_type VARCHAR NOT NULL DEFAULT 'in_app',
    ADD billing_type VARCHAR NOT NULL DEFAULT 'nurl',
    ADD auction_second_price BOOLEAN NOT NULL DEFAULT FALSE,
    ADD currency VARCHAR NOT NULL DEFAULT 'RUB';

-- Связывание таблицы ssp и валют
ALTER TABLE ssp.ssps ADD CONSTRAINT ssps_currencies_fk FOREIGN KEY (currency) REFERENCES ssp.currencies(slug);

-- Связывание таблицы ssp и типов биллинга
ALTER TABLE ssp.ssps ADD CONSTRAINT ssps_billing_types_fk FOREIGN KEY (billing_type) REFERENCES ssp.billing_types(slug);


-- Создание таблицы типов источников трафика
CREATE TABLE ssp.source_traffic_types (
	slug varchar NOT NULL,
	CONSTRAINT source_traffic_types_pk PRIMARY KEY (slug)
);

-- Добавление данных в таблицу типов источников трафика
INSERT INTO ssp.source_traffic_types (slug) VALUES('in_app');
INSERT INTO ssp.source_traffic_types (slug) VALUES('mobile_web');
INSERT INTO ssp.source_traffic_types (slug) VALUES('desktop');
INSERT INTO ssp.source_traffic_types (slug) VALUES('smart');

-- Связывание таблицы ssp и source_traffic_types
ALTER TABLE ssp.ssps ADD CONSTRAINT ssps_source_traffic_types_fk FOREIGN KEY (source_traffic_type) REFERENCES ssp.source_traffic_types(slug);



-- Создание таблицы типов интеграций с SSP
CREATE TABLE ssp.integration_types (
	slug varchar NOT NULL,
	CONSTRAINT integration_types_pk PRIMARY KEY (slug)
);

-- Добавление данных в таблицу типов интеграций
INSERT INTO ssp.integration_types (slug) VALUES('ortb');
INSERT INTO ssp.integration_types (slug) VALUES('vast');
INSERT INTO ssp.integration_types (slug) VALUES('feed');

-- Связывание таблицы ssp и integration_types
ALTER TABLE ssp.ssps ADD CONSTRAINT ssps_integration_types_fk FOREIGN KEY (integration_type) REFERENCES ssp.integration_types(slug);



-- Создание таблицы типов форматов
CREATE TABLE ssp.format_types
(
    slug VARCHAR NOT NULL,
    CONSTRAINT format_types_pk PRIMARY KEY (slug)
);

-- Добавление данных в таблицу ssp.format_types
INSERT INTO ssp.format_types (slug)
VALUES
    ('banner'),
    ('video'),
    ('push'),
    ('clickunder'),
    ('native');



-- Создание связующей таблицы ssp и format_types
CREATE TABLE ssp.ssps_to_format_types
(
    format_type_slug VARCHAR NOT NULL,
    ssp_slug VARCHAR NOT NULL,
    CONSTRAINT ssps_to_format_types_pk PRIMARY KEY (format_type_slug, ssp_slug),
    CONSTRAINT ssps_to_format_types_format_types_fk FOREIGN KEY (format_type_slug) REFERENCES ssp.format_types (slug),
    CONSTRAINT ssps_to_format_types_ssps_fk FOREIGN KEY (ssp_slug) REFERENCES ssp.ssps (slug) ON DELETE CASCADE
);

-- Добавить в связующую таблицу "ssps_to_format_types" записи для всех SSP, у которых стоит supports_clickunder_format = TRUE.
INSERT INTO ssp.ssps_to_format_types (format_type_slug, ssp_slug)
SELECT 'clickunder', slug
FROM ssp.ssps
WHERE supports_clickunder_traffic = TRUE;




-- Удаление колонки "supports_clickunder_traffic"
ALTER TABLE ssp.ssps
    DROP COLUMN supports_clickunder_traffic;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ssp.ssps_to_format_types;
DROP TABLE ssp.format_types;
DROP TABLE ssp.ssps;
-- +goose StatementEnd
