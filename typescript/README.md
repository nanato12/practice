# TypeScript

## docker
```shell
$ docker build . -t typescript
$ docker run -itd --name ts-practice -p 8888:8000 -v $PWD/workspace:/workspace --workdir /workspace typescript
$ docker exec -it ts-practice /bin/sh
```

## init
初回に実行したコマンド

```shell
$ npm install --save-dev typescript
$ ./node_modules/.bin/tsc --init
$ npm install --save-dev webpack webpack-cli ts-loader
```
