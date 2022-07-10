CREATE DATABASE IF NOT EXISTS vulneb;
USE vulneb;

CREATE TABLE IF NOT EXISTS accounts(
  id VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL
);

INSERT INTO accounts (id, password) VALUES ('Charlotte', 'faF2ji5');
INSERT INTO accounts (id, password) VALUES ('Liam', 'K9egkEir');
INSERT INTO accounts (id, password) VALUES ('Adam', 'Pofiwwkf#1');
