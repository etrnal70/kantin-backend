DROP TABLE IF EXISTS kantin.user;
CREATE TABLE kantin.user(
  id serial NOT NULL,
  firstname varchar(20),
  lastname varchar(30),
  email varchar(30) NOT NULL UNIQUE,
  password varchar(64) NOT NULL,
  createdAt timestamp NOT NULL,
  PRIMARY KEY(id)
);
