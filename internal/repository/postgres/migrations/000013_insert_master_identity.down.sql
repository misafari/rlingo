DELETE FROM tenant_member
       WHERE tenant_id = (SELECT id FROM tenant WHERE slug = 'Master') AND
             user_id = (SELECT id FROM users WHERE email = 'ms.safari@outlook.com');
DELETE FROM tenant WHERE slug = 'Master';
DELETE FROM users WHERE email = 'ms.safari@outlook.com';