# go-programming

### Overview

Go programming code, tested and run on Docker containers.

### Commands

**Running Tests using Please Build**

```bash
# reflection directory
plz test //reflection:test -vvv

# unit-testing directory
plz test //unit-testing:license_plates_test -vvv
```

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

**Install Please Build locally (MacOS).**

```bash
curl https://get.please.build | bash
source ~/.profile
plz --version
plz --help
```

**Initialize Please in the Repo**

```bash
plz init
```

### Directories

| Directory Name          | Description                                                                                             |
|-------------------------|---------------------------------------------------------------------------------------------------------|
| `.run`                  | Run configurations to use in the GoLand IDE.                                                            |
| `advanced-concurrency`  | Go code for advanced concurrency techniques, based on Chapter 9 of The Go Programming Language.         |
| `base`                  | Base `Dockerfile` with common logic for an Alpine Linux Go environment.                                 |
| `composite-types`       | Go code dealing with the basics of composite types, based on Chapter 4 of The Go Programming Language.  |
| `functions`             | Go code dealing with the basics of functions, based on Chapter 5 of The Go Programming Language.        |
| `go-tool`               | Shell and Go code utilizing the `go` CLI tool, based on Chapters 9 & 10 of The Go Programming Language. |
| `goroutines`            | Go code dealing with goroutines and concurrency, based on Chapter 8 of The Go Programming Language.     |
| `interfaces`            | Go code dealing with the basics of interfaces, based on Chapter 7 of The Go Programming Language.       |
| `low-level-programming` | Go code for writing *unsafe* low-level programming, based on Chapter 13 of The Go Programming Language. |
| `methods`               | Go code dealing with the basics of methods, based on Chapter 6 of The Go Programming Language.          |
| `reflection`            | Go code for reflection, based on Chapter 12 of The Go Programming Language.                             |
| `unit-testing`          | Go unit test coding basics, based on Chapter 11 of The Go Programming Language.                         |
| `.plzconfig`            | Please Build configuration file for the repository.                                                     |
| `BUILD`                 | Please Build rules for the top level directory of the repository.                                       |

### Version History

**[v1.0.1](https://github.com/AJarombek/go-programming/tree/v1.0.1) - Goroutine Article Update**

> Release Date: Sep 5th, 2022

* Additional examples of goroutines and channels for use in an upcoming 
[article on goroutines](https://jarombek.com/blog/sep-10-2022-goroutines).

**[v1.0.0](https://github.com/AJarombek/go-programming/tree/v1.0.0) - Initial Version**

> Release Date: Aug 27th, 2022

* Go language examples based on the book, *The Go Programming Language*, by Donovan and Kernighan.

### Resources

1. [The Go Programming Language](https://www.gopl.io/)
2. [Please Build Go Intro](https://please.build/codelabs/go_intro/#0)
3. [Please Build Go Rules](https://please.build/lexicon.html#go)
4. [Please Build Go Sample App](https://github.com/thought-machine/please-codelabs/tree/main/getting_started_go)
