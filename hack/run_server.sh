#!/usr/bin/env bash
export CRED_SERVER_CERT_LOCATION=hack/server-cert.pem;
export CRED_SERVER_DATABASE_DSN=postgresql://postgres:postgres@localhost/credentials?sslmode=disable;
export CRED_SERVER_ENCRYPTION_SECRET=test2;
export CRED_SERVER_JWT_SECRET=test;
export CRED_SERVER_LOCAL_STORAGE_LOCATION=/tmp;
export CRED_SERVER_PRIVATE_KEY_LOCATION=hack/server-key.pem

bin/server/cred-server-linux