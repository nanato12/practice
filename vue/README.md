# Vue
VueでGitHub Pageを作ってみる

## docker
```bash
$ docker-compose build
$ docker-compose up -d
```

## exec
```bash
$ docker exec -it vue_practice sh
```

## serve
```
/workspace # cd practice
/workspace/practice # npm run serve

> github-page@0.1.0 serve
> vue-cli-service serve

 INFO  Starting development server...
98% after emitting CopyPlugin

 DONE  Compiled successfully in 10397ms                                                                                                             1:19:48 PM


  App running at:
  - Local:   http://localhost:8080/practice/

```

## build
```
/workspace # cd practice
/workspace/practice # npm run build

> practice@0.1.0 build
> vue-cli-service build


⠋  Building for production...
```

## Github-page
```
/workspace/practice # npm run build
/workspace/practice # exit
$ cp -r dist ../../../docs
```
