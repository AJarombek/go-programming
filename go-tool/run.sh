#!/bin/sh

# Shell script which runs Go code and 'go' CLI commands.
# Author: Andrew Jarombek
# Date: 8/4/2022

cd go-tool || echo "go-tool directory"

# Examine the current go environment
printf "Go Environment Details: \n\n"
go env

# Test all the Go files
printf "\nTest Without Race Detection: \n\n"
go test -v .

# Prove that the race detector will find a race condition in race_condition_test.go
printf "\nTest With Race Detection: \n\n"
touch race_detection.log
go test -run=TestRaceCondition -race -v . >> race_detection.log

echo "Race Detection Exit Code: $?"
printf "Race Detection Log: \n\n"
cat race_detection.log

# Build and run an executable file from Go code
cd current_times || exit

printf "\nBuild and Run Go Code as an Executable Binary: \n\n"
go build current_times
./current_times

# Equivalent to the two commands above
go run current_times

printf "\nDocumentation for the Current Times Package: \n\n"
go doc current_times

printf "\nCurrently Available Packages: \n\n"

# This command prints all packages:
# go list ...

# But its return value is long, so this command specifically prints
# available packages with the 'unicode' prefix
go list unicode...
