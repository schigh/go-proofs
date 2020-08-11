#!/usr/bin/env bash

{
  go version
  printf "\nSTANDARD LIBRARY\n"
  go test -v -bench BenchmarkRun -benchmem ./stdlib
  printf "\n\nJSONITERATOR USING STDLIB-COMPATIBLE CONFIG\n"
  go test -v -bench BenchmarkRun -benchmem ./jsoniterator_compat
  printf "\n\nJSONITERATOR USING FASTEST CONFIG\n"
  go test -v -bench BenchmarkRun -benchmem ./jsoniterator_fastest
} > results.txt
