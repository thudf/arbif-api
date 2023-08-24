-- Executar o mysql no docker:
-- docker exec -it arbif-mysql bash

-- Executar login no mysql:
-- mysql -u root -p

-- Mostrar todos os databases;
SHOW DATABASES;

------------------------------------------------
CREATE DATABASE IF NOT EXISTS db_arbif;
USE db_arbif;

DROP TABLE IF EXISTS leads;

CREATE TABLE leads(
  id VARCHAR(50) PRIMARY KEY,
  credit_value DECIMAL(8, 2) NOT NULL,
  total_installment INT NOT NULL,
  cnpj VARCHAR(14) NOT NULL,
  name VARCHAR(100) NOT NULL,
  cpf VARCHAR(11) NOT NULL,
  email VARCHAR(256) NOT NULL,
  phone VARCHAR(15) NOT NULL,
  opt_in BOOLEAN NOT NULL,
  created_at TIMESTAMP default CURRENT_TIMESTAMP NOT NULL
);