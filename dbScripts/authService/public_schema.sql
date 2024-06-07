CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA pg_catalog;

CREATE TABLE IF NOT EXISTS users {
  id uuid PRIMARY KEY DEFAULT uuid_generate_v1mc(),
  email TEXT NOT NULL UNIQUE,
  first_name TEXT,
  last_name TEXT,
  user_password TEXT NOT NULL,
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
}

ALTER TABLE users OWNER TO postgres