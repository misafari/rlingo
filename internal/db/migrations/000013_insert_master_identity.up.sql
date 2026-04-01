INSERT INTO tenant(slug, name, plan, status, created_at, updated_at)
VALUES ('Master', 'Master Tenant', 'PRO', 'ACTIVE', NOW(), NOW());

INSERT INTO users(email, full_name, password_hash)
VALUES ('ms.safari@outlook.com', 'Essi Safari', 'hashed_password');

INSERT INTO tenant_member(tenant_id, user_id, role, invited_by, joined_at)
VALUES ((SELECT id FROM tenant WHERE slug = 'Master'),
        (SELECT id FROM users WHERE email = 'ms.safari@outlook.com'),
        'ADMIN',
        (SELECT id FROM users WHERE email = 'ms.safari@outlook.com'),
        now())