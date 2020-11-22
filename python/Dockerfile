FROM python:3.9-alpine

RUN apk update
RUN apk add gcc musl-dev

RUN pip install --upgrade pip
RUN pip install --upgrade setuptools
RUN pip install flake8
RUN pip install black

WORKDIR /workspace
