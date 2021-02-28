#! /bin/bash

docker build . -t golang
docker run -itd --name go-practice -p 8000:3000 -v $PWD/workspace:/go/src/workspace golang
docker exec -it go-practice /bin/sh
