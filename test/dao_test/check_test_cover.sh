#!/bin/bash
go test -covermode=count -coverprofile=coverprofile.cov -run="^Test" -coverpkg=$(go list ./... | grep -v "/test" | tr n ,) ./...
go tool cover -html=coverprofile.cov
rm ./coverprofile.cov
