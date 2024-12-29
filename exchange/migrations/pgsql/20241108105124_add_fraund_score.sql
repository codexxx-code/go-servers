-- +goose Up
-- +goose StatementBegin

-- Добавление столбца fraud_score в таблицу ssp.ssps
ALTER TABLE ssp.ssps
    ADD fraud_score VARCHAR NOT NULL DEFAULT 'disable';

-- Создание таблицы fraud_score
CREATE TABLE ssp.fraud_score
(
    slug varchar NOT NULL,
    CONSTRAINT fraud_score_pk PRIMARY KEY (slug)
);

-- Добавление данных в таблицу fraud_score
INSERT INTO ssp.fraud_score (slug)
VALUES ('disable'),
       ('real_time'),
       ('pre_bid');

-- Связывание таблицы ssp и fraud_score
ALTER TABLE ssp.ssps
    ADD CONSTRAINT ssps_fraud_score_pk FOREIGN KEY (fraud_score) REFERENCES ssp.fraud_score (slug);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ssp.fraud_score;
-- +goose StatementEnd