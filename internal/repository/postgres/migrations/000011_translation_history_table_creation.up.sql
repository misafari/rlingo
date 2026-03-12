CREATE TABLE translation_history
(
    id             UUID PRIMARY KEY     DEFAULT uuid_generate_v4(),
    translation_id UUID        NOT NULL REFERENCES translation,
    old_value      TEXT,
    new_value      TEXT,
    changed_by     UUID        NOT NULL REFERENCES users,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
