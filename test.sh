#!/bin/sh

run_test_suite() {
    arg=$1

    echo "Running $arg unit tests"
    rm -rf ./monorepo/$arg/go.mod ./monorepo/$arg/go.sum
    go test ./monorepo/$arg/tests/unit/... -v
}

run_test_suite "antifraud"