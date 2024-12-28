-- +goose Up
-- +goose StatementBegin

CREATE TABLE zodiac.prompts (
	id int NOT NULL,
	"case" varchar NOT NULL,
	"text" varchar NOT NULL,
	CONSTRAINT promts_pk PRIMARY KEY (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS zodiac.prompts;
-- +goose StatementEnd
