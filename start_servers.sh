#!/bin/bash
go run ./cmd/server/main.go -storage=postgres &
go run ./cmd/client/main.go