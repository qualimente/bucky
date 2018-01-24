# Overview #

Bucky is a simple program to generate logs for the purpose of testing log-shipping sub-systems.

[![Build Status](https://travis-ci.org/qualimente/bucky.svg?branch=master)](https://travis-ci.org/qualimente/bucky)

# Developing #

1. Clean up build env `make clean`
2. Build the software `make build test`
3. Run the software `bin/darwin/amd64/bucky generate`

# Building Releases #

1. Clean up build env `make clean`
2. Build the software `make all`

Releases will be created in the `releases` directory

# Running #

`bucky` prints bucky messages to standard out with one message printed in json format every second, by default.

options can be printed with `--help`, `bucky generate`'s are most interesting permitting control over period, count, and format (text, json):

```
NAME:
   bucky generate - Run the program

USAGE:
   bucky generate [command options] [arguments...]

OPTIONS:
   --format value  Specify the output format: text, json
   --period value  Specify the period between starting log message emitters. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h' (default: 1s)
   --count value   Specify the count of messages emitted per period
```

# Etymology #

The name Bucky comes from [bucking](https://en.wikipedia.org/wiki/Log_bucking), which is the process of rough cutting logs.
