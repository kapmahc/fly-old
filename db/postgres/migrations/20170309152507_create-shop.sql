-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE shop_stores (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  address VARCHAR(255) NOT NULL,
  manager VARCHAR(255) NOT NULL,
  tel VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE shop_catalogs (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  parent_id BIGINT,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE shop_vendors (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE shop_products (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  vendor_id BIGINT REFERENCES shop_vendors,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE shop_products_catalogs (
  catalog_id BIGINT REFERENCES shop_catalogs,
  product_id     BIGINT REFERENCES shop_products,
  PRIMARY KEY (product_id, catalog_id)
);

CREATE TABLE shop_variants(
  id BIGSERIAL PRIMARY KEY,
  sku VARCHAR(64) NOT NULL,
  product_id BIGINT REFERENCES shop_products,
  price NUMERIC(12,2) NOT NULL,
  cost NUMERIC(12,2) NOT NULL,
  weight NUMERIC(12,2) NOT NULL,
  height NUMERIC(12,2) NOT NULL,
  width NUMERIC(12,2) NOT NULL,
  length NUMERIC(12,2) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_shop_variants_sku ON shop_variants (sku);


CREATE TABLE shop_journals (
  id BIGSERIAL PRIMARY KEY,
  action VARCHAR(255) NOT NULL,
  quantity BIGINT NOT NULL,
  store_id BIGINT REFERENCES shop_stores,
  variant_id  BIGINT REFERENCES shop_variants,
  user_id BIGINT REFERENCES users,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE shop_stocks (
  id BIGSERIAL PRIMARY KEY,
  quantity BIGINT NOT NULL,
  store_id BIGINT REFERENCES shop_stores,
  variant_id  BIGINT REFERENCES shop_variants,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE shop_properties (
  id BIGSERIAL PRIMARY KEY,
  key VARCHAR(255) NOT NULL,
  val VARCHAR(2048) NOT NULL,
  variant_id BIGINT REFERENCES shop_variants,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_shop_properties_key_variant ON shop_properties (key, variant_id);
CREATE INDEX idx_shop_properties_key ON shop_properties (key);


CREATE TABLE shop_addresses (
  id BIGSERIAL PRIMARY KEY,
  first_name VARCHAR(32) NOT NULL,
  middle_name VARCHAR(32) NOT NULL,
  last_name VARCHAR(32) NOT NULL,
  zip VARCHAR(12) NOT NULL,
  apt VARCHAR(16) NOT NULL,
  street VARCHAR(255) NOT NULL,
  city VARCHAR(32) NOT NULL,
  state VARCHAR(32) NOT NULL,
  country VARCHAR(32) NOT NULL,
  phone VARCHAR(32) NOT NULL,
  user_id BIGINT REFERENCES users,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_shop_addresses_first_name ON shop_addresses(first_name);
CREATE INDEX idx_shop_addresses_last_name ON shop_addresses(last_name);
CREATE INDEX idx_shop_addresses_zip ON shop_addresses(zip);
CREATE INDEX idx_shop_addresses_city ON shop_addresses(city);
CREATE INDEX idx_shop_addresses_state ON shop_addresses(state);
CREATE INDEX idx_shop_addresses_country ON shop_addresses(country);
CREATE INDEX idx_shop_addresses_phone ON shop_addresses(phone);

CREATE TABLE shop_orders(
  id BIGSERIAL PRIMARY KEY,
  number VARCHAR(128) NOT NULL,
  item_total NUMERIC(12,2) NOT NULL,
  adjustment_total NUMERIC(12,2) NOT NULL,
  payment_total NUMERIC(12,2) NOT NULL,
  total NUMERIC(12,2) NOT NULL,
  state VARCHAR(16) NOT NULL,
  shipment_state VARCHAR(16) NOT NULL,
  payment_state VARCHAR(16) NOT NULL,
  user_id BIGINT REFERENCES users,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_shop_orders_number ON shop_orders (number);
CREATE INDEX idx_shop_orders_state ON shop_orders (state);
CREATE INDEX idx_shop_orders_payment_state ON shop_orders (payment_state);
CREATE INDEX idx_shop_orders_shipment_state ON shop_orders (shipment_state);

CREATE TABLE shop_line_items(
  id BIGSERIAL PRIMARY KEY,
  price NUMERIC(12,2) NOT NULL,
  quantity INT NOT NULL,
  variant_id BIGINT REFERENCES shop_variants,
  order_id BIGINT REFERENCES shop_orders,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE shop_payment_methods(
  id BIGSERIAL PRIMARY KEY,
  type VARCHAR(16) NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  active BOOLEAN NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_shop_payment_methods_type ON shop_payment_methods (type);
CREATE INDEX idx_shop_payment_methods_name ON shop_payment_methods (name);

CREATE TABLE shop_payments(
  id BIGSERIAL PRIMARY KEY,
  amount NUMERIC(12,2) NOT NULL,
  order_id BIGINT REFERENCES shop_orders,
  payment_method_id BIGINT REFERENCES shop_payment_methods,
  state VARCHAR(16) NOT NULL,
  response_code VARCHAR(32),
  avs_response TEXT,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_shop_payments_state ON shop_payments(state);

CREATE TABLE shop_zones(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  active BOOLEAN NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_shop_zones_name ON shop_zones(name);

CREATE TABLE shop_countries(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_shop_countries_name ON shop_countries(name);

CREATE TABLE shop_states(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  country_id BIGINT REFERENCES shop_countries,
  zone_id BIGINT REFERENCES shop_zones,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_shop_states_country_name ON shop_states(country_id, name);

CREATE TABLE shop_shipping_methods(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  tracking VARCHAR(255) NOT NULL,
  logo VARCHAR(255) NOT NULL,
  active BOOLEAN NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX shop_shipping_methods_name ON shop_shipping_methods(name);

CREATE TABLE shop_zones_shipping_methods (
  shipping_method_id BIGINT REFERENCES shop_shipping_methods,
  zone_id     BIGINT REFERENCES shop_zones,
  PRIMARY KEY (shipping_method_id, zone_id)
);

CREATE TABLE shop_shipments(
  id BIGSERIAL PRIMARY KEY,
  number VARCHAR(255) NOT NULL,
  tracking VARCHAR(255) NOT NULL,
  cost NUMERIC(12,2) NOT NULL,
  shipped_at TIMESTAMP WITHOUT TIME ZONE,
  state VARCHAR(16) NOT NULL,
  shipping_method_id BIGINT REFERENCES shop_shipping_methods,
  order_id BIGINT REFERENCES shop_orders,
  address_id BIGINT REFERENCES shop_addresses,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX shop_shipments_number ON shop_shipments(number);
CREATE INDEX shop_shipments_state ON shop_shipments(state);

CREATE TABLE shop_return_authorizations(
  id BIGSERIAL PRIMARY KEY,
  number VARCHAR(255) NOT NULL,
  state VARCHAR(16) NOT NULL,
  amount NUMERIC(12,2) NOT NULL,
  reason TEXT NOT NULL,
  order_id BIGINT REFERENCES shop_orders,
  enter_by_id BIGINT REFERENCES users,
  enter_at TIMESTAMP WITHOUT TIME ZONE,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX shop_return_authorizations_number ON shop_return_authorizations(number);
CREATE INDEX shop_return_authorizations_state ON shop_return_authorizations(state);

CREATE TABLE shop_return_inventory_units(
  id BIGSERIAL PRIMARY KEY,
  state VARCHAR(16) NOT NULL,
  quantity INT NOT NULL,
  variant_id BIGINT REFERENCES shop_variants,
  order_id BIGINT REFERENCES shop_orders,
  shipment_id BIGINT REFERENCES shop_shipments,
  return_authorization_id BIGINT REFERENCES shop_return_authorizations,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX shop_return_inventory_units_state ON shop_return_inventory_units(state);

CREATE TABLE shop_chargebacks(
  id BIGSERIAL PRIMARY KEY,
  state VARCHAR(16) NOT NULL,
  amount NUMERIC(12,2) NOT NULL,
  order_id BIGINT REFERENCES shop_orders,
  operator_id BIGINT REFERENCES users,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX shop_chargebacks_state ON shop_chargebacks(state);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE shop_chargebacks;
DROP TABLE shop_return_inventory_units;
DROP TABLE shop_return_authorizations;
DROP TABLE shop_shipments;
DROP TABLE shop_zones_shipping_methods;
DROP TABLE shop_shipping_methods;
DROP TABLE shop_states;
DROP TABLE shop_countries;
DROP TABLE shop_zones;
DROP TABLE shop_payments;
DROP TABLE shop_payment_methods;
DROP TABLE shop_line_items;
DROP TABLE shop_orders;
DROP TABLE shop_addresses;
DROP TABLE shop_properties;
DROP TABLE shop_stocks;
DROP TABLE shop_journals;
DROP TABLE shop_variants;
DROP TABLE shop_products_catalogs;
DROP TABLE shop_products;
DROP TABLE shop_vendors;
DROP TABLE shop_catalogs;
DROP TABLE shop_stores;
