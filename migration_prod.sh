#!/bin/bash
goose -dir ./migrations postgres "host=pg-prod port=5432 dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} sslmode=disable" up -v