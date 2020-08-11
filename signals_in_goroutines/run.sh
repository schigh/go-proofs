#!/usr/bin/env bash

case $1 in
  main)
    go run ./sig_main/main.go
    ;;
  dead)
    go run ./sig_goroutine_dead/main.go
    ;;
  alive)
    go run ./sig_goroutine_running/main.go
    ;;
  none)
    go run ./sig_none/main.go
    ;;
  *)
    echo "usage: ./run.sh (main|dead|alive|none) ... see ./README.md"
    ;;
  esac
