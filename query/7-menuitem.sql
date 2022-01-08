DROP TABLE IF EXISTS kantin.menuitem;
CREATE TABLE kantin.menuitem(
  id serial NOT NULL,
  seller_id int,
  name varchar(50) NOT NULL UNIQUE,
  description varchar(120),
  category int NOT NULL,
  price int NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_seller FOREIGN KEY(seller_id) REFERENCES kantin.seller(id),
  CONSTRAINT fk_category FOREIGN KEY(category) REFERENCES kantin.category(id)
);
