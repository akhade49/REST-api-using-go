CREATE TABLE products
(
    id bigint NOT NULL,
    name text COLLATE pg_catalog."default",
    CONSTRAINT product_pkey PRIMARY KEY (id)
    
);

INSERT INTO products (id, name) VALUES (1,'TV');
INSERT INTO products (id, name) VALUES (2,'AC');