DROP TABLE IF EXISTS kantin.order;
CREATE TABLE kantin.order(
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL,
  seller_id uuid NOT NULL,
  status text NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES kantin.user(id),
  CONSTRAINT fk_seller FOREIGN KEY(seller_id) REFERENCES kantin.seller(id)
);
