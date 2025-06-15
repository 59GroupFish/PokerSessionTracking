-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sessions (
    id UUID PRIMARY KEY,
    bb_value REAL NOT NULL,
    cash_chip_ratio REAL NOT NULL,
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    inactive_date TIMESTAMP WITHOUT TIME ZONE,
    game_id UUID NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;
-- +goose StatementEnd
