CREATE TABLE IF NOT EXISTS translation
(
    id         TEXT PRIMARY KEY,
    key        TEXT NOT NULL,
    locale     TEXT NOT NULL,
    text       TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);