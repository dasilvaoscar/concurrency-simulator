#!/bin/sh

echo "Running antifraud unit tests"
go test ./antifraud/tests/unit/... -v
