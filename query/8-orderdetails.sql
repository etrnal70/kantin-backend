DROP TABLE IF EXISTS kantin.orderdetails;
CREATE TABLE kantin.orderdetails(
  id serial NOT NULL,
  menuitem_id int NOT NULL,
  quantity int NOT NULL,
  comment varchar(120),
  PRIMARY KEY(id),
  CONSTRAINT fk_menuitem FOREIGN KEY(menuitem_id) REFERENCES kantin.menuitem(id)
);
