#!/bin/sh

# run gosec on all files, abort if failure found
gosec -severity high ./...

# run devtool validate to validate the service
# devtool validate (Go Kit Validate does not do anything yet)
