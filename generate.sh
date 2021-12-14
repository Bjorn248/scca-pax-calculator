#!/bin/bash

echo "Building site generator..."
go build main.go
echo "Done"

echo "Running site generator..."
./main

if ! git diff-index --quiet HEAD --
then
  echo "Found a difference in git!"
  ./lint_and_fix.sh
else
  echo "Found no difference in the generated file, exiting"
  exit 0
fi
