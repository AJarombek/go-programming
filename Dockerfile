# Dockerfile for running all the Go tests in the repository on Alpine Linux using Go 1.18
# Author: Andrew Jarombek
# Date: 9/21/2022

FROM ajarombek/go-alpine-linux-programming:latest

LABEL maintainer="andrew@jarombek.com" \
      version="1.0.0" \
      description="Dockerfile for running all the Go tests in the repository on Alpine Linux using Go 1.18"

COPY . .
RUN chmod -R 777 .

ENTRYPOINT ["plz", "test", "//...", "-i", "test", "-vvv"]