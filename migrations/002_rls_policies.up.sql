-- Enable Row-Level Security for multi-tenancy
ALTER TABLE translations ENABLE ROW LEVEL SECURITY;
ALTER TABLE projects ENABLE ROW LEVEL SECURITY;

-- Function to set tenant context
CREATE OR REPLACE FUNCTION set_tenant(tenant uuid)
RETURNS void AS $$
BEGIN
  PERFORM set_config('app.tenant_id', tenant::text, true);
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;

-- Policy to isolate tenants
CREATE POLICY tenant_isolation_translations ON translations
USING (tenant_id = current_setting('app.tenant_id')::uuid);

CREATE POLICY tenant_isolation_projects ON projects
USING (tenant_id = current_setting('app.tenant_id')::uuid);
