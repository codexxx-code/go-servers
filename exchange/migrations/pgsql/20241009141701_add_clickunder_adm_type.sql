-- +goose Up
-- +goose StatementBegin

-- Создание таблицы с типами adm ответов для кликандера
CREATE TABLE ssp.clickunder_adm_types (
                                          slug varchar NOT NULL,
                                          CONSTRAINT clickunder_adm_types_unique UNIQUE (slug)
);
COMMENT ON COLUMN ssp.clickunder_adm_types.slug IS 'Название типа adm ответа для кликандера';

-- Заполняем значениями
INSERT INTO ssp.clickunder_adm_types (slug) VALUES('url');
INSERT INTO ssp.clickunder_adm_types (slug) VALUES('xml');

-- Создаем поле в таблице ssps для хранения типа adm ответа для кликандера
ALTER TABLE ssp.ssps ADD clickunder_adm_type varchar NULL;
COMMENT ON COLUMN ssp.ssps.clickunder_adm_type IS 'Тип adm, который будет возвращаться для кликандер запроса';

-- Связываем поле с таблицей типов
ALTER TABLE ssp.ssps ADD CONSTRAINT ssps_clickunder_adm_types_fk FOREIGN KEY (clickunder_adm_type) REFERENCES ssp.clickunder_adm_types(slug);

-- Некоторые правки в рамках другой задачи
ALTER TABLE ssp.ssps RENAME COLUMN multiplication_factor TO clickunder_drum_size;
ALTER TABLE ssp.ssps ALTER COLUMN clickunder_drum_size DROP NOT NULL;
COMMENT ON COLUMN ssp.ssps.clickunder_drum_size IS 'Размер барабана для кликандер запроса';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE ssp.ssps RENAME COLUMN clickunder_drum_size TO multiplication_factor;
ALTER TABLE ssp.ssps DROP CONSTRAINT ssps_clickunder_adm_types_fk;
ALTER TABLE ssp.ssps DROP COLUMN clickunder_adm_type;
DROP TABLE ssp.clickunder_adm_types;
-- +goose StatementEnd
