-- +goose Up
-- +goose StatementBegin

-- Удаляем поле, связанное с языком из промптов
ALTER TABLE generator.prompts DROP COLUMN "language";

-- Удаляем связь с таблицей справочников
ALTER TABLE horoscope.horoscopes DROP CONSTRAINT forecasts_zodiacs_fk;

-- Создаем типы
CREATE TYPE horoscope.zodiac AS ENUM ('aries', 'taurus', 'gemini', 'cancer', 'leo', 'virgo', 'libra', 'scorpio', 'sagittarius', 'capricorn', 'aquarius', 'pisces');
CREATE TYPE horoscope.language AS ENUM ('russian', 'english');
CREATE TYPE horoscope.timeframe AS ENUM ('day', 'week', 'month', 'year');
CREATE TYPE horoscope.horoscope_type AS ENUM ('single', 'couple');

CREATE TYPE generator.prompt_case AS ENUM ('createHoroscope');

-- Очищаем всю таблицу промптов и гороскопов
DELETE FROM generator.prompts;
DELETE FROM horoscope.horoscopes;

-- Удаляем первичный ключ у таблицы промптов
ALTER TABLE generator.prompts DROP CONSTRAINT IF EXISTS prompts_pk;
ALTER TABLE generator.prompts DROP COLUMN IF EXISTS id;

-- Меняем поле case в таблице промптов на новый тип
ALTER TABLE generator.prompts ALTER COLUMN "case" TYPE generator.prompt_case USING "case"::generator.prompt_case;

-- Добавляем первичный ключ к таблице промптов
ALTER TABLE generator.prompts ADD PRIMARY KEY ("case");

-- Удаляем поле language из таблицы промптов
ALTER TABLE generator.prompts DROP COLUMN IF EXISTS "language";

-- Добавляем дефолтный промпт
INSERT INTO generator.prompts ("case", "text") VALUES ('createHoroscope', 'Создай пожалуйста гороскоп для любого знака зодиака');

-- Переименовываем и добавляем поля в таблицу гороскопов
ALTER TABLE horoscope.horoscopes RENAME COLUMN date TO date_from;
ALTER TABLE horoscope.horoscopes ADD COLUMN date_to DATE;
ALTER TABLE horoscope.horoscopes RENAME COLUMN zodiac TO primary_zodiac;
ALTER TABLE horoscope.horoscopes ALTER COLUMN primary_zodiac TYPE horoscope.zodiac USING primary_zodiac::horoscope.zodiac;
ALTER TABLE horoscope.horoscopes ADD COLUMN secondary_zodiac horoscope.zodiac;
ALTER TABLE horoscope.horoscopes ADD COLUMN timeframe horoscope.timeframe;
ALTER TABLE horoscope.horoscopes ADD COLUMN language horoscope.language;
ALTER TABLE horoscope.horoscopes ADD COLUMN horoscope_type horoscope.horoscope_type;

-- Удаляем лишнюю таблицу зодиаков
DROP TABLE horoscope.zodiacs;

-- Переименовываем таблицу промптов
ALTER TABLE generator.prompts RENAME TO prompt_templates;

-- Переименовываем поле text в таблице промптов
ALTER TABLE generator.prompt_templates RENAME COLUMN "text" TO "template";



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'This migration is irreversible';
-- +goose StatementEnd
