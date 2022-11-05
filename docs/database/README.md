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

| Serial | Name                     | Image              |  Units | Price | Discount | Annotations | Description |
|--------|--------------------------|--------------------|--------|-------|----------|-------------|-------------|
| 1      | Aceite esencial de clavo | /products/image/1  |  12ML  | 7.99  | 0        | 665,83 €/L  | El aceite...|

### Images collection:

| Serial | Image                          |
|--------|--------------------------------|
| 1      | https://i.ibb.co/jGc94N2/1.jpg |

