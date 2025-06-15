-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS players (
    id UUID PRIMARY KEY,
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    inactive_date TIMESTAMP WITHOUT TIME ZONE,
    game_id UUID NOT NULL,
    FOREIGN KEY (game_id) REFERENCES games(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE players;
-- +goose StatementEnd
