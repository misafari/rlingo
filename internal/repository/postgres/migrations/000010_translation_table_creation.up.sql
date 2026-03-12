CREATE TYPE translation_status AS ENUM ('DRAFT', 'REVIEW', 'APPROVED', 'OUTDATED');

CREATE TABLE translation
(
    id            UUID PRIMARY KEY            DEFAULT uuid_generate_v4(),
    key_id        UUID               NOT NULL REFERENCES translation_key,
    language_code VARCHAR(50)        NOT NULL REFERENCES language,
    value         TEXT,
    plural_forms  JSONB,
    status        translation_status NOT NULL DEFAULT 'DRAFT',
    translated_by UUID               NOT NULL REFERENCES users,
    reviewed_by   UUID REFERENCES users,
    reviewed_at   TIMESTAMPTZ,
    created_at    TIMESTAMPTZ        NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ        NOT NULL DEFAULT NOW()
);
