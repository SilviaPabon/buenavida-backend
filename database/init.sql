-- Database: buenavida

-- DROP DATABASE IF EXISTS buenavida;

CREATE DATABASE buenavida
    template 'template0'
    ENCODING = 'UTF8'
    LC_COLLATE = 'es_CO.UTF-8'
    LC_CTYPE = 'es_CO.UTF-8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

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
    -- 5 Numbers, 2 decimals
    total NUMERIC (5, 2) NOT NULL CHECK (total > 0),
    discount NUMERIC (5, 2) NOT NULL CHECK (discount > 0),
    CONSTRAINT fk_orders_users
	FOREIGN KEY ("idUser")
	    REFERENCES users (id)
	    ON UPDATE CASCADE
	    ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS "shoppingCart";

CREATE TABLE IF NOT EXISTS "cart"
(
    "idUser" INTEGER NOT NULL,
    "idArticle" CHAR ( 24 ) NOT NULL,
    amount SMALLINT NOT NULL CHECK (amount > 0),
    CONSTRAINT fk_shoppingCart_users
	FOREIGN KEY ("idUser")
	    REFERENCES users (id)
	    ON UPDATE CASCADE
	    ON DELETE CASCADE,
    CONSTRAINT cart_check_product_id
	check ("idArticle" ~ '^[0-9a-fA-F]{24}$')
);

-- DROP TABLE IF EXISTS favorites;

CREATE TABLE IF NOT EXISTS favorites
(
    "idUser" INTEGER NOT NULL,
    "idArticle" CHAR ( 24 ) NOT NULL,
    CONSTRAINT fk_favorites_users
	FOREIGN KEY ("idUser")
	    REFERENCES users (id)
	    ON UPDATE CASCADE
	    ON DELETE CASCADE,
    CONSTRAINT favorites_check_product_id
	check ("idArticle" ~ '^[0-9a-fA-F]{24}$')
);
