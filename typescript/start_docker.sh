#! /bin/bash

docker build . -t typescript
docker run -itd --name ts-practice -p 8080:8080 -v $PWD/workspace:/workspace --workdir /workspace typescript
docker exec -it ts-practice /bin/sh
