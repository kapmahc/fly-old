-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE shop_vendors (
  ID SERIAL PRIMARY KEY,
  description TEXT NOT NULL,
  vendor_id BIGINT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE shop_products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(32) NOT NULL,
  description TEXT NOT NULL,
  vendor_id BIGINT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE shop_variants(
  id SERIAL PRIMARY KEY,
  sku VARCHAR(64) NOT NULL,
  product_id BIGINT NOT NULL,
  price NUMBER(12,2) NOT NULL,
  cost NUMBER(12,2) NOT NULL,
  weight NUMBER(12,2) NOT NULL,
  height NUMBER(12,2) NOT NULL,
  width NUMBER(12,2) NOT NULL,
  length NUMBER(12,2) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_shop_variants_sky ON shop_variants (sku);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE t1;
