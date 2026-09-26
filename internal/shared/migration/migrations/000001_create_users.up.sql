CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    tg_id BIGINT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_active_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

CREATE INDEX IF NOT EXISTS idx_users_last_active_at ON users (last_active_at);
CREATE INDEX IF NOT EXISTS idx_users_tg_id ON users (tg_id);
CREATE INDEX IF NOT EXISTS idx_users_party_id ON users (party_id);


ALTER TABLE users
    ADD CONSTRAINT phone_format
        CHECK (phone ~ '^\+[1-9][0-9]{6,14}$') NOT VALID;
