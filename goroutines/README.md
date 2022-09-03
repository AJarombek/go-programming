### Overview

Go code related to goroutines, based on Chapter 8 of *The Go Programming Language*.

### Files

| Filename                          | Description                                                                                         |
|-----------------------------------|-----------------------------------------------------------------------------------------------------|
| `basic_example`                   | Basics example program for using goroutines.                                                        |
| `basics_test.go`                  | Basics of goroutines and concurrency in Go.                                                         |
| `buffered_channels_test.go`       | Achieving concurrency with buffered channels in Go.                                                 |
| `cancel_goroutine_test.go`        | Canceling goroutines using channels in Go.                                                          |
| `channels_test.go`                | Basics of channels which communicate and send data between goroutines in Go.                        |
| `Dockerfile`                      | Dockerfile for testing the `goroutines` program.                                                    |
| `exercises.json`                  | JSON file containing a list of exercises.  Used as input data for the goroutine pipeline example.   |
| `go.mod`                          | Go module file for the `goroutines` test code.                                                      |
| `limit_goroutines_test.go`        | Techniques for limiting the number of concurrent goroutines.                                        |
| `multiplexing_select_test.go`     | Using the `select` statement for multiplexing (receiving and handling data from multiple channels). |
| `pipelines_test.go`               | Using goroutines to form a pipeline in Go.                                                          |
| `unidirectional_channels_test.go` | Building upon `pipelines_test.go`, except with unidirectional channels.                             |
| `wait_group_test.go`              | Using a `WaitGroup` to run a variable number of operations in parallel with goroutines.             |