DROP POLICY IF EXISTS tenant_isolation_translations ON translations;
DROP POLICY IF EXISTS tenant_isolation_projects ON projects;
DROP FUNCTION IF EXISTS set_tenant(uuid);
