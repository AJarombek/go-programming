# Dockerfile for Go programming on Alpine Linux
# Author: Andrew Jarombek
# Date: 7/1/2022

FROM golang:1.18.3-alpine3.16

LABEL maintainer="andrew@jarombek.com" \
      version="1.0.0" \
      description="Dockerfile for Go programming on Alpine Linux"

RUN apk add gcc libc-dev linux-headers curl bash

RUN curl https://get.please.build > please.sh && \
    bash ./please.sh

ENV PATH="${PATH}:/root/.please/bin"

STOPSIGNAL SIGTERM