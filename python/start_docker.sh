#! /bin/bash

docker-compose build
docker-compose up -d
docker exec -it python_practice sh
