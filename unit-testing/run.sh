#!/bin/sh

# Shell script which runs Go code and 'go' CLI commands.
# Author: Andrew Jarombek
# Date: 8/9/2022

set -x

# Test all the Go files
go test -v .

# Test two specific test functions related to country codes
go test -v -run="TestValidCountryCode|TestInvalidCountryCode"