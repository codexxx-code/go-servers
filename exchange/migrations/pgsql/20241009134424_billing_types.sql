-- +goose Up
-- +goose StatementBegin

-- Создаем таблицу типов биллинга
CREATE TABLE ssp.billing_types
(
    slug varchar NOT NULL,
    CONSTRAINT billing_types_pk PRIMARY KEY (slug)
);

-- Добавляем данные в таблицу billing_types
INSERT INTO ssp.billing_types (slug) VALUES('nurl');
INSERT INTO ssp.billing_types (slug) VALUES('burl');
INSERT INTO ssp.billing_types (slug) VALUES('adm');
INSERT INTO ssp.billing_types (slug) VALUES('imp');
INSERT INTO ssp.billing_types (slug) VALUES('nurl_and_burl');

-- Связываем таблицу dsps и billing_types
ALTER TABLE ssp.dsps ADD CONSTRAINT dsps_billing_types_fk FOREIGN KEY (billing_url_type) REFERENCES ssp.billing_types(slug);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

ALTER TABLE ssp.dsps DROP CONSTRAINT dsps_billing_types_fk;
DROP TABLE ssp.billing_types;

-- +goose StatementEnd
