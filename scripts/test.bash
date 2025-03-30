#!/bin/bash
# Exit on first error
set -e

# Navigate to the parent directory where the project is
cd "$(dirname "$0")/.."

# Define Go versions to test
go_versions=("1.24.0" "1.23.0" "1.22.0" "1.21.0" "1.20" "1.19" "1.18" "1.17" "1.16" "1.15" "1.14" "1.13" "1.12")

# Install and test with different Go versions
for version in "${go_versions[@]}"
do
    echo "Setting up go$version"
    go install golang.org/dl/go$version@latest
    go$version download

    echo "Testing with go$version"
    # Clear cache and run test
    go$version clean -testcache
    go$version test ./... -v -race -covermode=atomic -timeout 1h
done
