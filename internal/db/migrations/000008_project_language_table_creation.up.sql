CREATE TABLE project_language
(
    project_id    UUID        NOT NULL REFERENCES project,
    language_code VARCHAR(50) NOT NULL REFERENCES language,
    is_base       BOOLEAN     NOT NULL DEFAULT FALSE,
    enabled_at    TIMESTAMPTZ          DEFAULT NOW()
);