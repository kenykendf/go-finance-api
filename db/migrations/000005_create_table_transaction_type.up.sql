CREATE TABLE IF NOT EXISTS type_transaction (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(25) NOT NULL,
    description VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);