-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS persons (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	surname VARCHAR(255) NOT NULL,
	patronymic VARCHAR(255),
	age INTEGER,
	gender VARCHAR(50),
	nationality VARCHAR(50),
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE persons
-- +goose StatementEnd
