#!/bin/bash

# This script is used to build the TranslateServer application.
# It compiles the Go code, runs tests, and generates coverage reports.

set -e

buildDir="build/out"
binaryDir="$buildDir/TranslateServer"

go build -o $binaryDir cmd/TranslateServer/main.go

testPath="$buildDir/unit-test"
coveragePath="$testPath/coverage.out"

mkdir -p $testPath
go test -v -shuffle on ./... -coverprofile=$coveragePath
go tool cover -func=$coveragePath -o $testPath/coverage.txt
go tool cover -html=$coveragePath -o $testPath/coverage.html
