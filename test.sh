#!/bin/sh

echo "Running antifraud unit tests"
go test ./monorepo/antifraud/tests/unit/... -v
