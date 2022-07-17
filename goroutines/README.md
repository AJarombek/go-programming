### Overview

Go code related to goroutines, based on Chapter 8 of *The Go Programming Language*.

### Files

| Filename                          | Description                                                                                       |
|-----------------------------------|---------------------------------------------------------------------------------------------------|
| `basics_test.go`                  | Basics of goroutines and concurrency in Go.                                                       |
| `channels_test.go`                | Basics of channels which communicate and send data between goroutines in Go.                      |
| `Dockerfile`                      | Dockerfile for testing the `goroutines` program.                                                  |
| `exercises.json`                  | JSON file containing a list of exercises.  Used as input data for the goroutine pipeline example. |
| `go.mod`                          | Go module file for the `goroutines` test code.                                                    |
| `pipelines_test.go`               | Using goroutines to form a pipeline in Go.                                                        |
| `unidirectional_channels_test.go` | Building upon `pipelines_test.go`, except with unidirectional channels.                           |