DROP TABLE IF EXISTS kantin.order;
CREATE TABLE kantin.order(
  id serial NOT NULL,
  user_id int NOT NULL,
  status varchar(15) NOT NULL,
  createdAt timestamp NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES kantin.user(id)
);
