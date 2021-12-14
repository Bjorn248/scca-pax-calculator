#!/bin/bash

echo "Building site generator..."
go build main.go
echo "Done"

echo "Running site generator..."
./main

branch_name="generated-changes"

if ! git diff-index --quiet HEAD --
then
  echo "Found a difference in git!"
  git stash
  echo "Checking for existence of $branch_name branch..."
  if git rev-parse --verify $branch_name
  then
    echo "Found branch, deleting"
    git branch -D $branch_name
  fi
  git checkout -b $branch_name
  git stash pop
  ./lint_and_fix.sh
else
  echo "Found no difference in the generated file, exiting"
  exit 0
fi
