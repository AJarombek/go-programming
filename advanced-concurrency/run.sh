#!/bin/sh

# Shell script which runs Go code and 'go' CLI commands.
# Author: Andrew Jarombek
# Date: 8/4/2022

# Test all the Go files
go test -v .

# Test with the race detector.  The race detector is an
# automated way to try and detect race conditions in concurrent code.
go test -race -v .