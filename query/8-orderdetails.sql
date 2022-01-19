DROP TABLE IF EXISTS kantin.orderdetails;
CREATE TABLE kantin.orderdetails(
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  order_id uuid NOT NULL,
  menuitem_id uuid NOT NULL,
  quantity int NOT NULL,
  comment text,
  PRIMARY KEY(id),
  CONSTRAINT fk_menuitem FOREIGN KEY(menuitem_id) REFERENCES kantin.menuitem(id),
  CONSTRAINT fk_order FOREIGN KEY(order_id) REFERENCES kantin.order(id)
);
