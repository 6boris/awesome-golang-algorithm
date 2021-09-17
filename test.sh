#!/bin/bash
#!/usr/bin/env bash

set -e
echo "" > coverage.txt

function test_lcof(){
    test_dir="lcof"
    echo "Test Dir ${test_dir}"
    go clean -testcache
    go test -coverprofile=profile.out -covermode=atomic -parallel 16 -p 16 ./$test_dir/*
}
function test_leetcode(){
  go clean -testcache
  test_dir="leetcode"
  echo "Test Dir ${test_dir}"
  for d in $(ls $test_dir); do
    {
     echo $d
      go test -coverprofile=profile.out -covermode=atomic -parallel 16 -p 16 ./$test_dir/$d/*
      if [ -f profile.out ]; then
          cat profile.out >> coverage.txt
          rm profile.out
      fi
    } &
  done
  wait
}

test_lcof
test_leetcode
