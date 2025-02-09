#!/bin/bash

go test -race -cover -coverprofile=coverage.out ./...

grep -v -E -f .covignore coverage.out > coverage.filtered.out
mv coverage.filtered.out coverage.out

go tool cover -html=coverage.out
