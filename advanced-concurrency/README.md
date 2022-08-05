### Overview

Go code related to composite types, based on Chapter 4 of *The Go Programming Language*.

### Files

| Filename                     | Description                                                             |
|------------------------------|-------------------------------------------------------------------------|
| `Dockerfile`                 | Dockerfile for testing the `advanced-concurrency` program.              |
| `go.mod`                     | Go module file for the `advanced-concurrency` test code.                |
| `monitor_goroutine_test.go`  | Go code for using the *monitor goroutine* pattern in concurrent code.   |
| `mutex_test.go`              | Go code for using binary semaphores and mutex locks in concurrent code. |
| `run.sh`                     | Shell script which runs on the Docker container upon startup.           |
| `rw_mutex_test.go`           | Go code for using read/write mutex locks in concurrent code.            |
| `serial_confinement_test.go` | Go code for using the *serial confinement* pattern in concurrent code.  |