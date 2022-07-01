# go-programming

### Overview

Go programming code, tested and run on Docker containers.

### Commands

**Create Base Docker Image**

```bash
docker login --username=ajarombek
docker image build -t go-alpine-linux-programming:latest ./base

# Push image to DockerHub with tag 'latest'
docker image tag go-alpine-linux-programming:latest ajarombek/go-alpine-linux-programming:latest
docker push ajarombek/go-alpine-linux-programming:latest
```

**Fix the `go.mod` file.**

```bash
go mod tidy
```

**Format the Go code.**

```bash
go fmt
```

### Directories

| Directory Name    | Description                                                                                            |
|-------------------|--------------------------------------------------------------------------------------------------------|
| `.run`            | Run configurations to use in the GoLand IDE.                                                           |
| `base`            | Base `Dockerfile` with common logic for an Alpine Linux Go environment.                                |
| `composite-types` | Go code dealing with the basics of composite types, based on Chapter 4 of The Go Programming Language. |

### Resources

1. [The Go Programming Language](https://www.gopl.io/)