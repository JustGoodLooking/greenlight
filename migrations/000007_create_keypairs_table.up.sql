CREATE TABLE IF NOT EXISTS keypairs (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    alias text NOT NULL,
    key_type text NOT NULL,
    public_key text NOT NULL UNIQUE,
    address text NOT NULL UNIQUE,S
    status text NOT NULL
);

CREATE INDEX idx_keypairs_user_id ON keypairs(user_id);
CREATE INDEX idx_keypairs_address ON keypairs(address);