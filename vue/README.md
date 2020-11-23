# Vue
VueでGitHub Pageを作ってみる

## docker
```bash
$ docker-compose build
$ docker-compose up -d
```

## exec
```
$ docker exec -it vue_practice sh


/workspace # cd practice
/workspace/practice # npm run build

> practice@0.1.0 build
> vue-cli-service build


⠋  Building for production...
/workspace/practice # cp -f docs ../
```

## run
```
/workspace/practice # npm run build
/workspace/practice # exit
$ cp -r dist ../../../docs
```
