#!/bin/bash

echo Iniciando ambiente local

docker network create go_crud_dynamodb

docker compose -f ./dynamodb/docker-compose.yml up -d 
