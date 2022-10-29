-- Database: buenavida

-- DROP DATABASE IF EXISTS buenavida;

CREATE DATABASE buenavida
    ENCODING = 'UTF8'
    LC_COLLATE = 'es_CO.UTF-8'
    LC_CTYPE = 'es_CO.UTF-8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

-- This command must be placed manually
\c buenavida;

-- DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users
(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR ( 125 ) NOT NULL,
    lastname VARCHAR ( 125 ) NOT NULL,
    mail VARCHAR ( 250 ) UNIQUE NOT NULL,
    password VARCHAR ( 250 ) NOT NULL
);

-- DROP TABLE IF EXISTS orders;

CREATE TABLE IF NOT EXISTS orders
(
    "idOrder" SERIAL NOT NULL PRIMARY KEY,
    "idUser" INTEGER NOT NULL,
    "orderDetails" JSON NOT NULL,
    total INTEGER NOT NULL,
    discount INTEGER NOT NULL,
	CONSTRAINT fk_orders_users
		FOREIGN KEY ("idUser")
			REFERENCES users (id)
			ON UPDATE CASCADE
			ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS "shoppingCart";

CREATE TABLE IF NOT EXISTS "shoppingCart"
(
    "idUser" INTEGER NOT NULL,
    "idArticle" INTEGER NOT NULL,
    amount INTEGER NOT NULL
	CONSTRAINT fk_shoppingCart_users
		FOREIGN KEY ("idUser")
			REFERENCES users (id)
			ON UPDATE CASCADE
			ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS favorites;

CREATE TABLE IF NOT EXISTS favorites
(
    "idUser" INTEGER NOT NULL,
    "idArticles" INTEGER NOT NULL
	CONSTRAINT fk_favorites_users
		FOREIGN KEY ("idUser")
			REFERENCES users (id)
			ON UPDATE CASCADE
			ON DELETE CASCADE
);
