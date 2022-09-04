CREATE table product (
     id serial PRIMARY KEY,
     uuid VARCHAR(255) NOT NULL,
     name VARCHAR (200) NOT NULL,
     price INT NOT NULL,
     active TINYINT(1),
     on_stock TINYINT(1)
);