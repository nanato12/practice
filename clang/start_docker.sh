#! /bin/bash

docker-compose build
docker-compose up -d
docker exec -it clang_practice sh
