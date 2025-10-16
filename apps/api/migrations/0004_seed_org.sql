-- 0004_seed_org.sql
-- Seed a default organization for testing

INSERT INTO org (id, name)
VALUES
  (1, 'Default Organization')
ON CONFLICT (id) DO NOTHING;
