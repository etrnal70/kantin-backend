DROP TABLE IF EXISTS kantin.seller;
CREATE TABLE kantin.seller(
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  name text NOT NULL,
  email text NOT NULL UNIQUE,
  password text NOT NULL,
  logo text,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(id)
);
