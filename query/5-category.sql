DROP TABLE IF EXISTS kantin.category;
CREATE TABLE kantin.category(
  id serial NOT NULL,
  value varchar(15) NOT NULL,
  PRIMARY KEY(id)
);

INSERT INTO kantin.category(value) VALUES ('FOOD');
INSERT INTO kantin.category(value) VALUES ('SNACK');
INSERT INTO kantin.category(value) VALUES ('BEVERAGE');
INSERT INTO kantin.category(value) VALUES ('DESSERT');
