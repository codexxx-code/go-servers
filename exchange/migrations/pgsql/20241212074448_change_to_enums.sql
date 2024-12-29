-- +goose Up
-- +goose StatementBegin

-- Удаляем связки между таблицами справочников и таблицей dsp
ALTER TABLE ssp.dsps
DROP constraint if exists dsps_billing_types_fk,
DROP constraint if exists dsps_integration_types_fk,
DROP constraint if exists dsps_source_traffic_types_fk;

-- Удаляем связки между таблицами справочников и таблицей ssp
ALTER TABLE ssp.ssps
DROP constraint if exists ssps_billing_types_fk,
DROP constraint if exists ssps_fraud_score_pk,
DROP constraint if exists ssps_integration_types_fk,
DROP constraint if exists ssps_source_traffic_types_fk;

-- Удаляем таблицы справочников
drop table if exists ssp.fraud_score;
DROP table if exists ssp.billing_types;
DROP table if exists ssp.integration_types;
DROP table if exists ssp.source_traffic_types;

DROP type if exists source_traffic_type;
DROP type if exists integration_type;
DROP type if exists format_type;
DROP type if exists billing_type;
DROP type if exists fraud_score;

-- Создаем типы данных
CREATE TYPE ssp.source_traffic_type AS ENUM ('in_app', 'mobile_web', 'desktop', 'smart');
CREATE TYPE ssp.integration_type AS ENUM ('ortb', 'vast', 'feed');
CREATE TYPE ssp.format_type AS ENUM ('banner', 'video', 'push', 'clickunder', 'native');
CREATE TYPE ssp.billing_type AS ENUM ('nurl', 'burl', 'adm', 'imp', 'nurl_and_burl');
CREATE TYPE ssp.fraud_score AS ENUM ('disable', 'real_time', 'pre_bid');

alter table ssp.dsps
alter column billing_url_type drop default,
alter column source_traffic_type drop default,
alter column integration_type drop default;

alter table ssp.ssps
alter column billing_type drop default,
alter column integration_type drop default,
alter column source_traffic_type drop default,
alter column fraud_score drop default;

-- Конвертируем стринги в типы у уже существующих полей в таблице dsp
ALTER TABLE ssp.dsps
ALTER COLUMN billing_url_type TYPE ssp.billing_type USING billing_url_type::ssp.billing_type,
ALTER COLUMN integration_type TYPE ssp.integration_type USING integration_type::ssp.integration_type,
ALTER COLUMN source_traffic_type TYPE ssp.source_traffic_type USING source_traffic_type::ssp.source_traffic_type;

-- Конвертируем стринги в типы у уже существующих полей в таблице ssp
ALTER TABLE ssp.ssps
ALTER COLUMN billing_type TYPE ssp.billing_type USING billing_type::ssp.billing_type,
ALTER COLUMN integration_type TYPE ssp.integration_type USING integration_type::ssp.integration_type,
ALTER COLUMN source_traffic_type TYPE ssp.source_traffic_type USING source_traffic_type::ssp.source_traffic_type,
ALTER COLUMN fraud_score TYPE ssp.fraud_score USING fraud_score::ssp.fraud_score;

-- Добавляем новые поля массивов перечислений в таблицу dsp
ALTER TABLE ssp.dsps
ADD COLUMN format_types ssp.format_type[] default '{}'::ssp.format_type[],
ADD COLUMN source_traffic_types ssp.source_traffic_type[] default '{}'::ssp.source_traffic_type[];

-- Добавляем новые поля массивов перечислений в таблицу ssp
ALTER TABLE ssp.ssps
ADD COLUMN format_types ssp.format_type[] default '{}'::ssp.format_type[],
ADD COLUMN source_traffic_types ssp.source_traffic_type[] default '{}'::ssp.source_traffic_type[];

-- Добавляем в каждую dsp массив значений из соединительной таблицы
UPDATE ssp.dsps d
SET format_types = (
    SELECT ARRAY_AGG(format_type_slug::ssp.format_type)
    FROM ssp.dsps_to_format_types
    WHERE dsp_slug = d.slug
);

-- Добавляем в каждую ssp массив значений из соединительной таблицы
UPDATE ssp.ssps s
SET format_types = (
    SELECT ARRAY_AGG(format_type_slug::ssp.format_type)
    FROM ssp.ssps_to_format_types
    WHERE ssp_slug = s.slug
);

-- Конвертируем поле source_traffic_type в source_traffic_types для dsp
UPDATE ssp.dsps
SET source_traffic_types = ARRAY[source_traffic_type::ssp.source_traffic_type];

-- Конвертируем поле source_traffic_type в source_traffic_types для ssp
UPDATE ssp.ssps
SET source_traffic_types = ARRAY[source_traffic_type::ssp.source_traffic_type];

-- Удаляем старое поле source_traffic_type у dsp
alter table ssp.dsps
drop column source_traffic_type;

-- Удаляем старое поле source_traffic_type у ssp
alter table ssp.ssps
drop column source_traffic_type;

-- Дропаем соединительные таблицы и связанные с ними таблицы
DROP TABLE ssp.dsps_to_format_types;
DROP TABLE ssp.ssps_to_format_types;
DROP TABLE ssp.format_types;

-- Делаем поля массивов енамов not null
UPDATE ssp.dsps
SET source_traffic_types = '{}'::ssp.source_traffic_type[]
WHERE source_traffic_types IS NULL;

UPDATE ssp.dsps
SET format_types = '{}'::ssp.format_type[]
WHERE format_types IS NULL;

UPDATE ssp.ssps
SET source_traffic_types = '{}'::ssp.source_traffic_type[]
WHERE source_traffic_types IS NULL;

UPDATE ssp.ssps
SET format_types = '{}'::ssp.format_type[]
WHERE format_types IS NULL;

ALTER TABLE ssp.ssps ALTER COLUMN source_traffic_types SET NOT NULL;
ALTER TABLE ssp.ssps ALTER COLUMN format_types SET NOT NULL;
ALTER TABLE ssp.dsps ALTER COLUMN source_traffic_types SET NOT NULL;
ALTER TABLE ssp.dsps ALTER COLUMN format_types SET NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
