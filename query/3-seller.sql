DROP TABLE IF EXISTS kantin.seller;
CREATE TABLE kantin.seller(
  id serial NOT NULL,
  name varchar(50) NOT NULL,
  email varchar(50) NOT NULL UNIQUE,
  password varchar(64) NOT NULL,
  orderlist int[],
  createdAt timestamp NOT NULL,
  PRIMARY KEY(id)
);
