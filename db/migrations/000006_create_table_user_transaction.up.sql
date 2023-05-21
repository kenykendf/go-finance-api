CREATE TABLE IF NOT EXISTS user_transaction (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    "date" TIMESTAMPTZ,
    category BIGINT NOT NULL,
    "type" BIGINT NOT NULL,
    note TEXT NOT NULL,
    amount INT NOT NULL,
    currency BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY(category) REFERENCES category(id),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY("type") REFERENCES type_transaction(id),
    FOREIGN KEY(currency) REFERENCES currency(id)
);

CREATE INDEX idx_user_trans ON user_transaction (id, user_id, currency, category);