CREATE TABLE translation_key
(
    id          UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
    project_id  UUID         NOT NULL REFERENCES project,
    key_name    VARCHAR(500) NOT NULL,
    description TEXT,
    tag         VARCHAR(20)[],
    is_plural   BOOLEAN      NOT NULL DEFAULT FALSE,
    created_by  UUID         NOT NULL REFERENCES users,
    created_at  TIMESTAMPTZ           DEFAULT NOW(),
    updated_at  TIMESTAMPTZ           DEFAULT NOW()
);

CREATE UNIQUE INDEX unique_idx_translation_key_project_key ON translation_key (project_id, key_name);
CREATE INDEX idx_translation_key_project_tag ON translation_key (project_id, tag);
