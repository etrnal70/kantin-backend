DROP TABLE IF EXISTS kantin.status;
CREATE TABLE kantin.status(
  id serial NOT NULL,
  value text NOT NULL,
  PRIMARY KEY(id)
);

INSERT INTO kantin.status(value) VALUES ('WAITING_APPROVAL');
INSERT INTO kantin.status(value) VALUES ('PROCESSING');
INSERT INTO kantin.status(value) VALUES ('REQUEST_FOR_CANCELLATION');
INSERT INTO kantin.status(value) VALUES ('READY');
INSERT INTO kantin.status(value) VALUES ('DONE');
INSERT INTO kantin.status(value) VALUES ('CANCELLED');
