FROM node:15-alpine

ENV HOME /root

RUN npm i -g @vue/cli

RUN apk update && \
    apk add git

RUN cd / && \
    yarn add vue-template-compiler && \
    cd ./node_modules/vue-template-compiler/ && \
    sed -i -e "s/var vueVersion = .*/var vueVersion = undefined/" index.js

COPY .vuerc ${HOME}/.vuerc
