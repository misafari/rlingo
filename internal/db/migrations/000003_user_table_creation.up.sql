CREATE TYPE user_role AS ENUM ('VIEWER', 'EDITOR', 'ADMIN');
CREATE TYPE users_status AS ENUM ('ACTIVE', 'SUSPENDED', 'DELETED');

CREATE TABLE users
(
    id            UUID PRIMARY KEY             DEFAULT uuid_generate_v4(),
    email         VARCHAR(255) UNIQUE NOT NULL,
    is_sso        BOOLEAN             NOT NULL DEFAULT FALSE,
    full_name     VARCHAR(255),
    password_hash VARCHAR(255),
    status        users_status        NOT NULL DEFAULT 'ACTIVE',
    last_login_at TIMESTAMPTZ                  DEFAULT NOW(),
    created_at    TIMESTAMPTZ                  DEFAULT NOW(),

    CONSTRAINT password_logic CHECK (
        (is_sso = FALSE AND password_hash IS NOT NULL) OR
        (is_sso = TRUE AND password_hash IS NULL)
    )
);