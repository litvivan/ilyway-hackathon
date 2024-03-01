
-- +migrate Up
CREATE TABLE items (
  id                 SERIAL PRIMARY KEY,
  title              varchar(255) NOT NULL,
  description        text NOT NULL,
  participant_count  int    NOT NULL,
  activity_type      varchar(255) NOT NULL,
  city               varchar(255) NOT NULL,
  author_name        varchar(255) NOT NULL,
  author_rating      float NOT NULL,
  image_url          varchar(255) NOT NULL,
  full_address       varchar(1000) NOT NULL,
  has_reward         boolean NOT NULL,
  duration           varchar(255) NOT NULL,
  start_at           timestamptz NOT NULL,
  created_at         timestamptz default (now() at time zone 'utc'),
  updated_at         timestamptz default (now() at time zone 'utc')
);

-- +migrate Down
DROP TABLE items;