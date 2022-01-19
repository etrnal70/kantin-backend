DROP TABLE IF EXISTS kantin.user;
CREATE TABLE kantin.user(
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  firstname text NOT NULL,
  lastname text,
  email text NOT NULL UNIQUE,
  password text NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(id)
);
