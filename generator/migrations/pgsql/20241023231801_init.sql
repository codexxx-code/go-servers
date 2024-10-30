-- +goose Up
-- +goose StatementBegin

CREATE SCHEMA zodiac;

CREATE TABLE zodiac.zodiacs (
	slug varchar NOT NULL,
	"name" varchar NOT NULL,
	CONSTRAINT zodiacs_pk PRIMARY KEY (slug)
);

INSERT INTO zodiac.zodiacs (slug, "name") VALUES('aries', 'Овен');
INSERT INTO zodiac.zodiacs (slug, "name") VALUES('taurus', 'Телец');
INSERT INTO zodiac.zodiacs (slug, "name") VALUES('gemini', 'Близнецы');
INSERT INTO zodiac.zodiacs (slug, "name") VALUES('cancer', 'Рак');
INSERT INTO zodiac.zodiacs (slug, "name") VALUES('leo', 'Лев');
INSERT INTO zodiac.zodiacs (slug, "name") VALUES('virgo', 'Дева');
INSERT INTO zodiac.zodiacs (slug, "name") VALUES('libra', 'Весы');
INSERT INTO zodiac.zodiacs (slug, "name") VALUES('scorpio', 'Скорпион');
INSERT INTO zodiac.zodiacs (slug, "name") VALUES('sagittarius', 'Стрелец');
INSERT INTO zodiac.zodiacs (slug, "name") VALUES('capricorn', 'Козерог');
INSERT INTO zodiac.zodiacs (slug, "name") VALUES('aquarius', 'Водолей');
INSERT INTO zodiac.zodiacs (slug, "name") VALUES('pisces', 'Рыбы');

CREATE TABLE zodiac.forecasts (
	id int GENERATED ALWAYS AS IDENTITY NOT NULL,
	"date" date NOT NULL,
	zodiac varchar NOT NULL,
	"text" varchar NOT NULL,
	CONSTRAINT forecasts_unique UNIQUE (id),
	CONSTRAINT forecasts_zodiacs_fk FOREIGN KEY (zodiac) REFERENCES zodiac.zodiacs(slug)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA zodiac;
-- +goose StatementEnd
