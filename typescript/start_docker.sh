#! /bin/bash

docker build . -t typescript
docker run -itd --name ts-practice -p 8888:8000 -v $PWD/workspace:/workspace --workdir /workspace typescript
docker exec -it ts-practice /bin/sh
