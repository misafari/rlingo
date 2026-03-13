CREATE TYPE invitation_status AS ENUM ('PENDING', 'ACCEPTED', 'EXPIRED', 'REVOKED');

CREATE TABLE invitation
(
    id            UUID PRIMARY KEY             DEFAULT uuid_generate_v4(),
    project_id    UUID                NOT NULL REFERENCES translation,
    tenant_id     UUID                NOT NULL REFERENCES tenant,
    invitee_email VARCHAR(250)        NOT NULL,
    invitee_role  user_role           NOT NULL,
    token         VARCHAR(255) UNIQUE NOT NULL,
    status        invitation_status   NOT NULL DEFAULT 'PENDING',
    expires_at    TIMESTAMPTZ         NOT NULL,
    invited_by    UUID                NOT NULL REFERENCES users,
    created_at    TIMESTAMPTZ         NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_invitation_project_id_status ON invitation (project_id, status);
CREATE INDEX idx_invitation_token ON invitation (token);
