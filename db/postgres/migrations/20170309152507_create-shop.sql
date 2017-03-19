-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE shop_addresses (
  ID SERIAL PRIMARY KEY,
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
  user_id BIGINT NOT NULL,
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
  id SERIAL PRIMARY KEY,
  number VARCHAR(128) NOT NULL,
  item_total NUMERIC(12,2) NOT NULL,
  adjustment_total NUMERIC(12,2) NOT NULL,
  payment_total NUMERIC(12,2) NOT NULL,
  total NUMERIC(12,2) NOT NULL,
  state VARCHAR(16) NOT NULL,
  shipment_state VARCHAR(16) NOT NULL,
  payment_state VARCHAR(16) NOT NULL,
  user_id BIGINT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_shop_orders_number ON shop_orders (number);
CREATE INDEX idx_shop_orders_state ON shop_orders (state);
CREATE INDEX idx_shop_orders_payment_state ON shop_orders (payment_state);
CREATE INDEX idx_shop_orders_shipment_state ON shop_orders (shipment_state);

CREATE TABLE shop_line_items(
  id SERIAL PRIMARY KEY,
  price NUMERIC(12,2) NOT NULL,
  quantity INT NOT NULL,
  variant_id BIGINT NOT NULL,
  order_id BIGINT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE shop_payment_methods(
  id SERIAL PRIMARY KEY,
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
  id SERIAL PRIMARY KEY,
  amount NUMERIC(12,2) NOT NULL,
  order_id BIGINT NOT NULL,
  payment_method_id BIGINT NOT NULL,
  state VARCHAR(16) NOT NULL,
  response_code VARCHAR(32),
  avs_response TEXT,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_shop_payments_state ON shop_payments(state);

CREATE TABLE shop_zones(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_shop_zones_name ON shop_zones(name);

CREATE TABLE shop_countries(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_shop_countries_name ON shop_countries(name);

CREATE TABLE shop_states(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  country_id BIGINT NOT NULL,
  zone_id BIGINT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX idx_shop_states_country_name ON shop_states(country_id, name);

CREATE TABLE shop_shipping_methods(
  id SERIAL PRIMARY KEY,
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
  shipping_method_id BIGINT NOT NULL,
  zone_id     BIGINT NOT NULL,
  PRIMARY KEY (shipping_method_id, zone_id)
);

CREATE TABLE shop_shipments(
  id SERIAL PRIMARY KEY,
  number VARCHAR(255) NOT NULL,
  tracking VARCHAR(255) NOT NULL,
  cost NUMERIC(12,2) NOT NULL,
  shipped_at TIMESTAMP WITHOUT TIME ZONE,
  state VARCHAR(16) NOT NULL,
  shipping_method_id BIGINT NOT NULL,
  order_id BIGINT NOT NULL,
  address_id BIGINT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX shop_shipments_number ON shop_shipments(number);
CREATE INDEX shop_shipments_state ON shop_shipments(state);

CREATE TABLE shop_return_authorizations(
  id SERIAL PRIMARY KEY,
  number VARCHAR(255) NOT NULL,
  state VARCHAR(16) NOT NULL,
  amount NUMERIC(12,2) NOT NULL,
  reason TEXT NOT NULL,
  order_id BIGINT NOT NULL,
  enter_by_id BIGINT NOT NULL,
  enter_at TIMESTAMP WITHOUT TIME ZONE,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX shop_return_authorizations_number ON shop_return_authorizations(number);
CREATE INDEX shop_return_authorizations_state ON shop_return_authorizations(state);

CREATE TABLE shop_return_inventory_units(
  id SERIAL PRIMARY KEY,
  state VARCHAR(16) NOT NULL,
  quantity INT NOT NULL,
  variant_id BIGINT NOT NULL,
  order_id BIGINT NOT NULL,
  shipment_id BIGINT NOT NULL,
  return_authorization_id BIGINT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX shop_return_inventory_units_state ON shop_return_inventory_units(state);

CREATE TABLE shop_chargebacks(
  id SERIAL PRIMARY KEY,
  state VARCHAR(16) NOT NULL,
  amount NUMERIC(12,2) NOT NULL,
  order_id BIGINT NOT NULL,
  operator_id BIGINT NOT NULL,
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
