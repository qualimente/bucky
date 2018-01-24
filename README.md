# Overview #

Bucky is a simple program to generate logs for the purpose of testing log-shipping sub-systems.

[![Build Status](https://travis-ci.org/qualimente/bucky.svg?branch=master)](https://travis-ci.org/qualimente/bucky)

# Developing #

1. Clean up build env `make clean`
2. Build the software `make build test`
3. Run the software `bin/darwin/amd64/bucky generate`
```
$ bin/darwin/amd64/bucky generate
{"appName":"bucky","level":"info","msg":"Running periodically","time":"2018-01-24T08:58:03-07:00"}
{"level":"info","msg":"error perspiciatis omnis. quia accusamus nam itaque magnam aut. culpa aut eveniet quia voluptates in voluptate dicta. repellat magni sequi et eveniet. pariatur laboriosam accusantium et. dolores id eveniet facere error corporis sunt. rem dolores error et. amet debitis repudiandae qui. ullam incidunt doloribus molestiae doloremque.","msgId":"1516809483476061000-0","time":"2018-01-24T08:58:03-07:00","type":"log"}
{"level":"info","msg":"numquam hic soluta dolor voluptas error! quia sed voluptatem. et fugiat exercitationem nemo voluptatem maiores eaque. quis commodi vel aut et est et dolores. qui optio et et quia ad.","msgId":"1516809484479212000-0","time":"2018-01-24T08:58:04-07:00","type":"log"}
{"level":"info","msg":"adipisci aliquid qui necessitatibus reiciendis error quis temporibus nesciunt. rerum veritatis laboriosam sit. sed nulla similique et dolor maiores sint. nemo ea adipisci magni non. adipisci animi non eius commodi delectus!","msgId":"1516809485483082000-0","time":"2018-01-24T08:58:05-07:00","type":"log"}
...
```

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
   --count value   Specify the count of messages emitted per period (default: 1)
```

By default, `bucky` will generate one json message per second.

The `count` parameter is the main tuning knob to use for increasing log message volume.
  
For example, to generate 100 messages per second, run:

```
bucky generate --count 100
```

# Etymology #

The name Bucky comes from [bucking](https://en.wikipedia.org/wiki/Log_bucking), which is the process of rough cutting logs.
