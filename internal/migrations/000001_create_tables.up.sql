BEGIN;

CREATE TABLE IF NOT EXISTS users(
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "username" VARCHAR(64) NOT NULL UNIQUE,
    "password" VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS data(
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "credentials_id" BIGINT,
    "banking_card_id" BIGINT,
    "metadata_id" BIGINT
);

CREATE TABLE IF NOT EXISTS banking_cards_data(
    "id" BIGSERIAL PRIMARY KEY,
    "data_id" BIGINT NOT NULL,
    "number" INTEGER,
    "valid_till" DATE,
    "cardholder_name" VARCHAR(128),
    "cvv" INTEGER
);

CREATE TABLE IF NOT EXISTS credentials_data(
   "id" BIGSERIAL PRIMARY KEY,
   "data_id" BIGINT NOT NULL,
   "login" VARCHAR(128),
   "password" VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS metadata(
   "id" BIGSERIAL PRIMARY KEY,
   "data_id" BIGINT NOT NULL,
   "meta" TEXT
);

COMMIT;
