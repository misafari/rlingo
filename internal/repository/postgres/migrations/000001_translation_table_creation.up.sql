CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE translation
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    key VARCHAR(50),
    locale VARCHAR(10),
    Text VARCHAR(250),
    created_at    TIMESTAMPTZ DEFAULT now(),
    updated_at    TIMESTAMPTZ DEFAULT now()
);