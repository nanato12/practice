FROM ubuntu:18.04

USER root
# SHELL ["/bin/bash", "-c"]

# ubuntu setup
ENV DEBCONF_NOWARNINGS yes

RUN apt-get update
# Cコンパイラを作成するのに必要なパッケージ
RUN apt-get install -y --no-install-recommends \
    locales \
    gcc \
    make \
    git \
    binutils \
    libc6-dev

# Set local
RUN locale-gen en_US.UTF-8
