CREATE TYPE tenant_plan AS ENUM ('FREE', 'PRO', 'ENTERPRISE');
CREATE TYPE tenant_status AS ENUM ('ACTIVE', 'SUSPENDED', 'DELETED');

CREATE TABLE tenant
(
    id         UUID PRIMARY KEY             DEFAULT uuid_generate_v4(),
    slug       VARCHAR(100) UNIQUE NOT NULL,
    name       VARCHAR(100)        NOT NULL,
    plan       tenant_plan         NOT NULL DEFAULT 'FREE',
    status     tenant_status       NOT NULL DEFAULT 'ACTIVE',
    created_at TIMESTAMPTZ                  DEFAULT NOW(),
    updated_at TIMESTAMPTZ                  DEFAULT NOW()
);