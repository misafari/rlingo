CREATE TABLE project_member
(
    project_id UUID      NOT NULL REFERENCES project,
    user_id    UUID      NOT NULL REFERENCES users,
    role       user_role NOT NULL DEFAULT 'VIEWER',
    invited_by UUID REFERENCES users,
    joined_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);