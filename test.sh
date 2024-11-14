#!/bin/sh

echo "Running antifraud unit tests"
rm -rf ./monorepo/antifraud/go.mod ./monorepo/antifraud/go.sum
go test ./monorepo/antifraud/tests/unit/... -v
