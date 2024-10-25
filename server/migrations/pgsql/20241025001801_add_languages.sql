-- +goose Up
-- +goose StatementBegin

CREATE TABLE zodiac.prompts_copy (
	id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
	"language" varchar NOT NULL,
	"case" varchar NOT NULL,
	"text" varchar NOT NULL,
	CONSTRAINT prompts_pk PRIMARY KEY (id)
);

INSERT INTO zodiac.prompts_copy ("case", "text", "language")
SELECT "case", "text", 'russian'
FROM zodiac.prompts;

DROP TABLE zodiac.prompts;

ALTER TABLE zodiac.prompts_copy RENAME TO prompts;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Так делать нельзя)))

-- +goose StatementEnd
