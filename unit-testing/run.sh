#!/bin/sh

# Shell script which runs Go code and 'go' CLI commands.
# Author: Andrew Jarombek
# Date: 8/9/2022

set -x

# Get all production code Go files in the current directory
go list -f={{.GoFiles}} .

# Get all test code Go files in the current directory
go list -f={{.TestGoFiles}} .

# Test all the Go files
go test -v .

# Test two specific test functions related to country codes
go test -v -run="TestValidCountryCode|TestInvalidCountryCode"

# Calculate code coverage on all the Go files
go test -v -coverprofile=c.out .

# Run benchmark Go tests
go test -bench=.

# Run benchmark Go tests along with memory allocation statistics
go test -bench=. -benchmem

# CPU profiling
go test -run=NONE -bench=. -cpuprofile=cpu.log .
go tool pprof -text -nodecount=10 ./unit-testing.test cpu.log

# Block (processes blocking goroutines) profiling
go test -run=NONE -bench=. -blockprofile=block.log .
go tool pprof -text -nodecount=10 ./unit-testing.test block.log

# Heap (Memory) profiling
go test -run=NONE -bench=. -memprofile=mem.log .
go tool pprof -text -nodecount=10 ./unit-testing.test mem.log
