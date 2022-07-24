### Overview

Go code related to goroutines, based on Chapter 8 of *The Go Programming Language*.

### Files

| Filename                          | Description                                                                                       |
|-----------------------------------|---------------------------------------------------------------------------------------------------|
| `basics_test.go`                  | Basics of goroutines and concurrency in Go.                                                       |
| `buffered_channels_test.go`       | Achieving concurrency with buffered channels in Go.                                               |
| `channels_test.go`                | Basics of channels which communicate and send data between goroutines in Go.                      |
| `Dockerfile`                      | Dockerfile for testing the `goroutines` program.                                                  |
| `exercises.json`                  | JSON file containing a list of exercises.  Used as input data for the goroutine pipeline example. |
| `go.mod`                          | Go module file for the `goroutines` test code.                                                    |
| `limit_goroutines_test.go`        | Techniques for limiting the number of concurrent goroutines.                                      |
| `pipelines_test.go`               | Using goroutines to form a pipeline in Go.                                                        |
| `unidirectional_channels_test.go` | Building upon `pipelines_test.go`, except with unidirectional channels.                           |
| `wait_group_test.go`              | Using a `WaitGroup` to run a variable number of operations in parallel with goroutines.           |