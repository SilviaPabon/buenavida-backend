-- Database: buenavida

-- *** *** *** *** *** ***
-- *** TABLES CREATION ***
-- *** *** *** *** *** ***

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

-- DROP TABLE IF EXISTS "orders"
CREATE TABLE IF NOT EXISTS orders
(
  "idOrder" SERIAL NOT NULL PRIMARY KEY, 
  "idUser" INTEGER NOT NULL,
  "date" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  CONSTRAINT fk_orders_users
    FOREIGN KEY ("idUser")
      REFERENCES users(id)
      ON UPDATE CASCADE
      ON DELETE CASCADE
);

-- DROP TABLE IF EXISTS "orders_has_products"
CREATE TABLE IF NOT EXISTS orders_has_products
(
  "idOrder" INTEGER NOT NULL, 
  "idArticle" CHAR ( 24 ) NOT NULL, 
  "amount" SMALLINT NOT NULL CHECK (amount > 0), 
  CONSTRAINT order_has_products_product_id
    check ("idArticle" ~ '^[0-9a-fA-F]{24}$'), 
  CONSTRAINT order_has_products_order_id
    FOREIGN KEY ("idOrder")
      REFERENCES orders ("idOrder")
      ON UPDATE CASCADE
      ON DELETE CASCADE
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

-- *** *** *** *** *** ***
-- *** Stored procedures ***
-- *** *** *** *** *** ***
create or replace procedure make_order(
	-- Reseive the user id as an argument
	user_id int 
)
language plpgsql
as $$
declare 
	cur cursor for select "idArticle", "amount" from cart where "idUser" = user_id;
	current_row record;
	order_id integer; 
	cart_length integer;
begin 
	-- Verify at least one product exists on cart
	select count("idUser") into cart_length from cart
	where "idUser" = user_id;
	
	if cart_length > 0 then
		-- Creates a new empty order
		insert into orders("idUser") values (user_id);
	
		-- Get the new order id
		select "idOrder" into order_id from orders 
		where "idUser" = user_id
		order by "date" desc limit 1;
		
		open cur;
		
		-- Loop for each cart row and insert into 
		-- the orders table
		loop
			-- Get the row value
			fetch cur into current_row;
			exit when not found;
		
			raise notice 'Product: %', current_row."idArticle";
		
			-- Insert the row into the orders table
			insert into orders_has_products  ("idOrder", "idArticle", "amount") 
			values (order_id, current_row."idArticle", current_row."amount");
		
		end loop;
		
		close cur;
	
		-- Clear the user cart
		delete from cart where "idUser" = user_id;
	end if;

	commit;
end; $$;
