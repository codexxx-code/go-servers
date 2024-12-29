-- +goose Up
-- +goose StatementBegin

-- Добавление столбцов в таблицу ssp.ssps
ALTER TABLE ssp.ssps
    ADD clickunder_adm_format VARCHAR NULL DEFAULT '';

-- Обновление значений clickunder_adm_format в зависимости от clickunder_adm_type
UPDATE ssp.ssps
SET clickunder_adm_format = '${ADM_URL}'
WHERE clickunder_adm_type = 'url';

-- Обновление значений clickunder_adm_format в зависимости от clickunder_adm_type
UPDATE ssp.ssps
SET clickunder_adm_format = '<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n<ad><popunderAd><url><![CDATA[${ADM_URL}]]></url></popunderAd></ad>'
WHERE clickunder_adm_type = 'xml';

-- Удаление внешнего ключа
ALTER TABLE ssp.ssps
DROP CONSTRAINT IF EXISTS ssps_clickunder_adm_types_fk;

-- Удаление столбца clickunder_adm_type из таблицы ssp.ssps
ALTER TABLE ssp.ssps
DROP COLUMN clickunder_adm_type;

--Удаление таблицы clickunder_adm_types
DROP TABLE ssp.clickunder_adm_types;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ssp.clickunder_adm_types;
-- +goose StatementEnd
