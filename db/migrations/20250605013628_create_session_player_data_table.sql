-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS session_player_data (
    id UUID PRIMARY KEY,
    buy_in_cash REAL NOT NULL,
    final_stack_chips REAL NOT NULL,
    session_id UUID NOT NULL,
    player_id UUID NOT NULL,
    FOREIGN KEY (session_id) REFERENCES sessions(id),
    FOREIGN KEY (player_id) REFERENCES players(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE session_player_data;
-- +goose StatementEnd