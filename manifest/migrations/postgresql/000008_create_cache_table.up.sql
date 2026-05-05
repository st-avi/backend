CREATE UNLOGGED TABLE cache
(
    key TEXT PRIMARY KEY,
    value JSONB NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_cache_expires_at ON cache (expires_at);