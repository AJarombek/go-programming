# Dockerfile for running the file_management program on Debian Linux
# Author: Andrew Jarombek
# Date: 6/14/2022

FROM golang:1.18.3-alpine3.16

LABEL maintainer="andrew@jarombek.com" \
      version="1.0.0" \
      description="Dockerfile for running the composite-type program on Alpine Linux using Go 1.18"

WORKDIR src
COPY . .

ENTRYPOINT ["go", "test", "main_test.go"]