-- +goose Up
-- +goose StatementBegin

CREATE SCHEMA ssp;

-- Добавили объект DSP
CREATE TABLE ssp.dsps
(
    slug                 VARCHAR NOT NULL,
    name                 VARCHAR NOT NULL,
    url                  VARCHAR NOT NULL,
    currency             VARCHAR NOT NULL,
    auction_second_price bool    NOT NULL,
    billing_url_type     VARCHAR NOT NULL
);

-- Добавили объект валюты
CREATE TABLE ssp.currencies
(
    slug   VARCHAR NOT NULL,
    name   VARCHAR NOT NULL,
    symbol VARCHAR NOT NULL,
    rate   DECIMAL NOT NULL
);

-- Добавили объект ssp
CREATE TABLE ssp.ssps
(
    slug                        VARCHAR NOT NULL,
    name                        VARCHAR NOT NULL,
    supports_clickunder_traffic bool    NOT NULL,
    timeout                     INTEGER NULL
);

-- Добавили объект настроек
CREATE TABLE ssp.settings
(
    margin                         DECIMAL NOT NULL,
    host                           VARCHAR NOT NULL,
    default_timeout                INTEGER NOT NULL,
    empty_second_price_reduce_coef DECIMAL NOT NULL,
    reduce_timeout_coef            DECIMAL NOT NULL,
    showcase_url                   VARCHAR NOT NULL
);

-- Выставили первичные ключи
ALTER TABLE ssp.dsps
    ADD CONSTRAINT dsps_pk PRIMARY KEY (slug);
ALTER TABLE ssp.currencies
    ADD CONSTRAINT currencies_pk PRIMARY KEY (slug);
ALTER TABLE ssp.ssps
    ADD CONSTRAINT ssps_pk PRIMARY KEY (slug);

-- Выставили связи между сущностями
ALTER TABLE ssp.dsps
    ADD CONSTRAINT dsps_currencies_fk FOREIGN KEY (currency) REFERENCES ssp.currencies (slug);


-- Заполняем данными
INSERT INTO ssp.settings (margin, host, default_timeout, empty_second_price_reduce_coef, reduce_timeout_coef, showcase_url)
VALUES (0.1, 'https://bid.sspnet.tech', 3000000, 0.1, 0.15, 'https://feedinform.online');

INSERT INTO ssp.currencies (slug, name, symbol, rate)
VALUES ('RUB', 'Российский рубль', '₽', 0.01);
INSERT INTO ssp.currencies (slug, name, symbol, rate)
VALUES ('USD', 'Доллар США', '$', 1);

-- Добавили комментарии к каждому полю таблиц
COMMENT ON COLUMN ssp.dsps.slug IS 'Системное имя';
COMMENT ON COLUMN ssp.dsps.name IS 'Человекочитаемое имя';
COMMENT ON COLUMN ssp.dsps.url IS 'URL, на который слать OpenRTB запросы';
COMMENT ON COLUMN ssp.dsps.currency IS 'Валюта, с которой мы торгуем';

COMMENT ON COLUMN ssp.currencies.slug IS 'Строковый идентификатор валюты';
COMMENT ON COLUMN ssp.currencies.name IS 'Название валюты';
COMMENT ON COLUMN ssp.currencies.symbol IS 'Символ валюты';
COMMENT ON COLUMN ssp.currencies.rate IS 'Курс валюты относительно доллара';

COMMENT ON COLUMN ssp.ssps.slug IS 'Строковый идентификатор';
COMMENT ON COLUMN ssp.ssps.name IS 'Человекочитаемое название';
COMMENT ON COLUMN ssp.ssps.supports_clickunder_traffic IS 'Поддерживает ли кликандер трафик';
COMMENT ON COLUMN ssp.ssps.timeout IS 'Таймаут ответа к SSP. В миллисекундах';

COMMENT ON COLUMN ssp.settings.margin IS 'Маржа';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Удаляем таблицы
DROP SCHEMA ssp;

-- +goose StatementEnd
