-- +goose Up
-- +goose StatementBegin

-- Добавляем связующую таблицу ssp и dsp
CREATE TABLE ssp.ssps_to_dsps (
	ssp_slug varchar NOT NULL,
	dsp_slug varchar NOT NULL,
	is_deleted boolean NOT NULL,
	CONSTRAINT ssps_to_dsps_pk PRIMARY KEY (ssp_slug,dsp_slug),
	CONSTRAINT ssps_to_dsps_ssps_fk FOREIGN KEY (ssp_slug) REFERENCES ssp.ssps(slug),
	CONSTRAINT ssps_to_dsps_dsps_fk FOREIGN KEY (dsp_slug) REFERENCES ssp.dsps(slug)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ssp.ssps_to_dsps;
-- +goose StatementEnd
