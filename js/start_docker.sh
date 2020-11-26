#! /bin/bash

docker-compose build
docker-compose up -d
docker exec -it js_practice sh
