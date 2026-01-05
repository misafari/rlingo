CREATE TABLE translation_key
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    project_id UUID NOT NULL,
    key VARCHAR(150) NOT NULL ,
    created_at    TIMESTAMPTZ DEFAULT now(),
    updated_at    TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX translation_key_project_id_tag_idx ON translation_key (project_id, key);
CREATE INDEX idx_translation_key_project_id ON translation_key(project_id);