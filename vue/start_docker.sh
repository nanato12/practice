#! /bin/bash

docker-compose build
docker-compose up -d
docker exec -it vue_practice sh
