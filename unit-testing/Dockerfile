# Dockerfile for running the unit testing program on Alpine Linux using Go 1.18
# Author: Andrew Jarombek
# Date: 8/8/2022

FROM ajarombek/go-alpine-linux-programming:latest

LABEL maintainer="andrew@jarombek.com" \
      version="1.0.0" \
      description="Dockerfile for running the unit testing program on Alpine Linux using Go 1.18"

WORKDIR src
COPY . .

CMD ["-c", ". ./run.sh"]
ENTRYPOINT ["/bin/sh"]