CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE tenants (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        TEXT NOT NULL,
    created_at  TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE projects (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id     UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    name          TEXT NOT NULL,
    default_locale TEXT NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE translations (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id   UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    project_id  UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    key         TEXT NOT NULL,
    locale      TEXT NOT NULL,
    text        TEXT,
    status      TEXT DEFAULT 'DRAFT' CHECK (status IN ('DRAFT','REVIEWED','APPROVED')),
    updated_at  TIMESTAMPTZ DEFAULT now()
);

CREATE INDEX idx_translations_project ON translations (tenant_id, project_id);
CREATE INDEX idx_translations_key_locale ON translations (tenant_id, key, locale);
