-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE erp_catalogs (
  ID SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  parent_id BIGINT,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE erp_vendors (
  ID SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE erp_products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  vendor_id BIGINT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE erp_products_catalogs (
  catalog_id BIGINT NOT NULL,
  product_id     BIGINT NOT NULL,
  PRIMARY KEY (product_id, catalog_id)
);

CREATE TABLE erp_variants(
  id SERIAL PRIMARY KEY,
  sku VARCHAR(64) NOT NULL,
  product_id BIGINT NOT NULL,
  price NUMERIC(12,2) NOT NULL,
  cost NUMERIC(12,2) NOT NULL,
  weight NUMERIC(12,2) NOT NULL,
  height NUMERIC(12,2) NOT NULL,
  width NUMERIC(12,2) NOT NULL,
  length NUMERIC(12,2) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_erp_variants_sku ON erp_variants (sku);

CREATE TABLE erp_properties (
  ID SERIAL PRIMARY KEY,
  key VARCHAR(255) NOT NULL,
  val VARCHAR(2048) NOT NULL,
  variant_id BIGINT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_erp_properties_key_variant ON erp_properties (key, variant_id);
CREATE INDEX idx_erp_properties_key ON erp_properties (key);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE erp_properties;
DROP TABLE erp_variants;
DROP TABLE erp_products_catalogs;
DROP TABLE erp_products;
DROP TABLE erp_vendors;
DROP TABLE erp_catalogs;
