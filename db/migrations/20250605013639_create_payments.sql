-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    payer VARCHAR(255) NOT NULL,
    payee VARCHAR(255) NOT NULL,
    paid_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    player_id UUID NOT NULL,
    session_id UUID NOT NULL,
    FOREIGN KEY (player_id) REFERENCES players(id),
    FOREIGN KEY (session_id) REFERENCES sessions(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payments;
-- +goose StatementEnd