DROP TABLE IF EXISTS kantin.menuitem;
CREATE TABLE kantin.menuitem(
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  seller_id uuid NOT NULL,
  name text NOT NULL UNIQUE,
  description text,
  category int NOT NULL,
  price int NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_seller FOREIGN KEY(seller_id) REFERENCES kantin.seller(id),
  CONSTRAINT fk_category FOREIGN KEY(category) REFERENCES kantin.category(id)
);
