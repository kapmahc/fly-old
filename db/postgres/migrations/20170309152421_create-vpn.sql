-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE vpn_users (
  id         BIGINT                      REFERENCES users,
  description TEXT NOT NULL,
  online     BOOLEAN                     NOT NULL DEFAULT FALSE,
  enable     BOOLEAN                     NOT NULL DEFAULT FALSE,
  start_up   DATE                        NOT NULL DEFAULT '2016-12-13',
  shut_down  DATE                        NOT NULL DEFAULT current_date,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE vpn_logs (
  id           BIGSERIAL PRIMARY KEY,
  user_id      BIGINT REFERENCES users,
  trusted_ip   INET,
  trusted_port SMALLINT,
  remote_ip    INET,
  remote_port  SMALLINT,
  start_up     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  shut_down    TIMESTAMP WITHOUT TIME ZONE,
  received     FLOAT                       NOT NULL DEFAULT '0.0',
  send         FLOAT                       NOT NULL DEFAULT '0.0'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE vpn_logs;
DROP TABLE vpn_users;
