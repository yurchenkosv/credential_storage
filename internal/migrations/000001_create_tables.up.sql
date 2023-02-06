BEGIN;

CREATE TABLE IF NOT EXISTS users(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "username" VARCHAR(64) NOT NULL UNIQUE,
    "password" VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS banking_cards_data(
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "number" VARCHAR(64),
    "valid_till" VARCHAR(64),
    "cardholder_name" VARCHAR(128),
    "cvv" VARCHAR(64)
);

CREATE TABLE IF NOT EXISTS credentials_data(
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL ,
    "login" VARCHAR(128),
    "password" VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS text_data(
    "id" SERIAL PRIMARY KEY ,
    "user_id" INTEGER NOT NULL,
    "data" TEXT
);

CREATE TABLE IF NOT EXISTS binary_data(
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "link" VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS metadata(
    "id" BIGSERIAL PRIMARY KEY,
    "data_id" INTEGER,
    "meta" TEXT
);

CREATE TABLE IF NOT EXISTS data(
    "id" SERIAL PRIMARY KEY ,
    "name" VARCHAR(128) NOT NULL,
    "user_id" INTEGER NOT NULL ,
    "banking_cards_data_id" INTEGER REFERENCES banking_cards_data(id) ON DELETE CASCADE,
    "credentials_data_id" INTEGER REFERENCES credentials_data(id) ON DELETE CASCADE,
    "text_data_id" INTEGER REFERENCES text_data(id) ON DELETE CASCADE,
    "binary_data_id" INTEGER REFERENCES binary_data ON DELETE CASCADE
);

COMMIT;
