# Flonfig
Turn your config into flags easily.
---
[![Build Status](https://travis-ci.org/JesusIslam/flonfig.svg?branch=master)](https://travis-ci.org/JesusIslam/flonfig)
[![Coverage Status](https://coveralls.io/repos/github/JesusIslam/flonfig/badge.svg?branch=master)](https://coveralls.io/github/JesusIslam/flonfig?branch=master)
[![GoDoc](https://godoc.org/github.com/JesusIslam/flonfig?status.svg)](https://godoc.org/github.com/JesusIslam/flonfig)
[![Go Report Card](https://goreportcard.com/badge/github.com/JesusIslam/flonfig)](https://goreportcard.com/report/github.com/JesusIslam/flonfig)

## Dependencies
- `go get github.com/BurntSushi/toml`

## What
Just like what the subtitle says, it is created so you (well, me actually) could turn a config file into flags automatically.
You can also add environment variable lookup to the config and flonfig will prioritize the value it got from there.
Supported types are:
- string
- int
- int64
- uint
- uint64
- float64
- duration
- duration_string (like `3h` or `20s`)

## Why
Because I want a flexible config flag and I don't want to define the flags hardcoded. Essentially, because I am a lazy bum.

## How
Just open the examples directory.

## TODO
- Test extracting from env
- Test real flag values

## Notes
Also, I would gladly accept PR.
