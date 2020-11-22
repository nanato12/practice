# Haskell

## docker
```bash
$ docker-compose build
$ docker-compose up -d
```

## exec
```
$ docker exec -it haskell_practice bash

root@haskell:/workspace# cd tutorial/hello/
root@haskell:/workspace/tutorial/hello# ghc main.hs
[1 of 1] Compiling Main             ( main.hs, main.o )
Linking main ...
root@haskell:/workspace/tutorial/hello# ./main
Hello World
```

## チュートリアル
- [tutorial](./workspace/tutorial)
