# Databases documentation

## Postgres (relational) database

```mermaid
erDiagram

USERS ||--o{ CART: have
USERS ||--o{ FAVORITES: have
USERS ||--|{ ORDERS: make
USERS {
  SERIAL  ID      PK    "NOT NULL"
  VARCHAR NAME          "NOT NULL"
  VARCHAR LASTNAME      "NOT NULL"
  VARCHAR EMAIL         "UNIQUE, NOT NULL"
  VARCHAR PASSWORD      "NOT NULL"
}

CART{
  INTEGER   IDUSER      FK  "NOT NULL"
  CHAR      IDARTICLE       "NOT NULL, CONSTRAINT CHECK"
  SMALLINT  AMOUNT          "NOT NULL, GREATER THAN 0"
}

ORDERS ||--|{ ORDERS_HAS_PRODUCTS: in
ORDERS{
  SERIAL    IDORDER PK  "NOT NULL"
  INTEGER   IDUSER  FK  "NOT NULL"
  TIMESTAMP DATE        "NOT NULL, DEFAULT NOW()"
}

ORDERS_HAS_PRODUCTS{
  INTEGER   IDORDER     FK  "NOT NULL"
  CHAR      IDARTICLE       "NOT NULL, CONSTRAINT CHECK"
  SMALLINT  AMOUNT          "NOT NULL, GREATER THAN 0"
  NUMERIC   PRICE           "NOT NULL, GREATER THAN 0"
  NUMERIC   DISCOUNT        "NOT NULL, GREATER THAN 0"
}

FAVORITES{
  INTEGER   IDUSER      FK    "NOT NULL"
  CHAR      IDARTICLE         "NOT NULL, CONSTRAINT CHECK"
}

```

## Mongo (non-relationsl) database: 


Here we store our products and it's images in two collections:

### Products collection: 

```json
{
  "serial": 1, 
  "name": "Aceite esencial de clavo", 
  "image": "/products/image/1",
  "units": "12ML",
  "price": 7.99, 
  "discount": 0, 
  "annotations": "665,83 €/L", 
  "description": "El aceite..."
}
```

### Images collection:

```json
  "serial": 1, 
  "image": "https://i.ibb.co/jGc94N2/1.jpg"
```

## Redis (in-memory) store

We use redis to store the valids refresh-tokens (as a whitelist): 

```json
{
  "user1@gmail.com": "74b1e197-b140-48f8-9094-784c52f72dc7"
  "user2@outlook.com": "54d96d51-9afd-43d4-8c88-7566b02d1e4f"
}
```

**Note:** As you can see, we don't store the token string, instead, we store the token UUID which is crated along with the token. At this way, we can verify if some provided token is authentic and is not expired (**each token has a 12 hours TTL**).
