CREATE TYPE project_status AS ENUM ('ACTIVE', 'ARCHIVED');

CREATE TABLE project
(
    id            UUID PRIMARY KEY        DEFAULT uuid_generate_v4(),
    tenant_id     UUID           NOT NULL REFERENCES tenant,
    name          VARCHAR(50)    NOT NULL,
    description   VARCHAR(100)   NOT NULL,
    status        project_status NOT NULL DEFAULT 'ACTIVE',
    created_by    UUID           NOT NULL REFERENCES users,
    created_at    TIMESTAMPTZ             DEFAULT now(),
    updated_at    TIMESTAMPTZ             DEFAULT now()
);

CREATE INDEX idx_project_tenant_id ON project (tenant_id);