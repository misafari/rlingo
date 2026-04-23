-- name: GetTenantIDByUserID :one
SELECT id FROM tenant AS t
    LEFT JOIN tenant_member AS tm
        ON t.id = tm.tenant_id
         WHERE tm.user_id = $1;

