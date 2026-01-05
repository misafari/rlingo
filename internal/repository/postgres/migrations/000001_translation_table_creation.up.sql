CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE translation
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    key_id     UUID NOT NULL,
    locale_id  UUID NOT NULL,
    Text       TEXT,
    created_at TIMESTAMPTZ      DEFAULT now(),
    updated_at TIMESTAMPTZ      DEFAULT now()
);

CREATE UNIQUE INDEX translation_key_locale_unique_idx ON translation (key_id, locale_id);
CREATE INDEX idx_translation_locale_id ON translation (locale_id);