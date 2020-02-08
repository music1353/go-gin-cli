CREATE TABLE auth (
  id serial PRIMARY KEY NOT NULL,
  account varchar DEFAULT '',
  password varchar DEFAULT '',
  role varchar DEFAULT 'member'
);

INSERT INTO auth (account, password, role) VALUES ('admin', '0000', 'admin');

