CREATE TABLE project
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50),
    created_at    TIMESTAMPTZ DEFAULT now(),
    updated_at    TIMESTAMPTZ DEFAULT now()
);