CREATE TABLE templates (
  id char(36) CONSTRAINT firstkey PRIMARY KEY,
  class varchar(40) NOT NULL,
  content varchar NOT NULL,
  created_at date NOT NULL,
  updated_at date
);