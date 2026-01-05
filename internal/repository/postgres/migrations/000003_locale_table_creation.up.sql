CREATE TABLE locale
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    project_id UUID NOT NULL,
    locale VARCHAR(35) NOT NULL ,
    is_default BOOLEAN NOT NULL DEFAULT FALSE,
    created_at    TIMESTAMPTZ DEFAULT now(),
    updated_at    TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX locale_project_id_tag_idx ON locale (project_id, locale);
CREATE INDEX idx_locale_project_id ON locale(project_id);