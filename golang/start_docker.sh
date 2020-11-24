#! /bin/bash

docker-compose build
docker-compose up -d
docker exec -it golang_practice sh
