#!/bin/bash

echo Encerrando ambiente local

docker compose -f ./dynamodb/docker-compose.yml down

docker network rm go_crud_dynamodb