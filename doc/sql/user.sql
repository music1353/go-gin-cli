CREATE TABLE users (
  id integer REFERENCES auth (id),
  name varchar DEFAULT '',
  phone varchar DEFAULT ''
);

INSERT INTO users (id, name, phone) VALUES ('1', '蘇靖軒', '0989000628');