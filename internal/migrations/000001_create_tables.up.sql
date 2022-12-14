BEGIN;

CREATE TABLE IF NOT EXISTS users(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "username" VARCHAR(64) NOT NULL UNIQUE,
    "password" VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS data(
    "id" SERIAL PRIMARY KEY ,
    "user_id" INTEGER NOT NULL ,
    "banking_cards_data_id" INTEGER,
    "credentials_data_id" INTEGER,
    "text_data_id" INTEGER,
    "binary_data_id" INTEGER
);

CREATE TABLE IF NOT EXISTS banking_cards_data(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(128) NOT NULL,
    "user_id" INTEGER NOT NULL ,
    "number" INTEGER,
    "valid_till" DATE,
    "cardholder_name" VARCHAR(128),
    "cvv" INTEGER
);

CREATE TABLE IF NOT EXISTS credentials_data(
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL ,
    "name" VARCHAR(128) NOT NULL,
    "login" VARCHAR(128),
    "password" VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS text_data(
    "id" SERIAL PRIMARY KEY ,
    "user_id" INTEGER NOT NULL,
    "name" VARCHAR(128) NOT NULL,
    "data" TEXT
);

CREATE TABLE IF NOT EXISTS binary_data(
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "name" VARCHAR(128) NOT NULL,
    "link" VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS metadata(
    "id" BIGSERIAL PRIMARY KEY,
    "data_id" BIGINT,
    "meta" TEXT
);

COMMIT;
