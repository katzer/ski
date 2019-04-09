# ski - Sascha Knows It <br>[![GitHub release](https://img.shields.io/github/release/appplant/ski.svg)](https://github.com/appplant/ski/releases) [![Build Status](https://travis-ci.com/appplant/ski.svg?branch=master)](https://travis-ci.com/appplant/ski) [![Build status](https://ci.appveyor.com/api/projects/status/f5imsl77fmg2omba/branch/master?svg=true)](https://ci.appveyor.com/project/katzer/goo/branch/master) [![Maintainability](https://api.codeclimate.com/v1/badges/e5995227dd52c2f7221e/maintainability)](https://codeclimate.com/github/appplant/ski/maintainability)

Execute commands or collect informations on multiple servers in parallel.

    $ ski -h

    Usage: ski [options...] matchers...
    Options:
    -c, --command   Execute command and return result
    -s, --script    Execute script and return result
    -t, --template  Template to be used to transform the output
    -j, --job       Execute job specified in file
    -n, --no-color  Print errors without colors
    -p, --pretty    Pretty print output as a table
    -w, --width     Width of output column in characters
    -h, --help      This help text
    -v, --version   Show version number

## Prerequisites

You'll need to add `ORBIT_HOME`, `ORBIT_KEY` and `ORBIT_BIN` first to your profile:

    $ export ORBIT_HOME=/path/to/orbit

## Installation

Download the latest version from the [release page][releases] and add the executable to your `PATH`.

## Usage

Execute shell commands:

    $ ski -c 'echo Greetings from $PACKAGE_NAME' mars pluto

    Greetings from Mars
    Greetings from Pluto

Execute shell scripts:

    $ ski -s greet.sh mars pluto

Execute SQL commands:

    $ ski -c 'SELECT * FROM DUAL' db

    D
    -
    X

Execute SQL scripts:

    $ ski -s dummy.sql db

Pretty table output:

    $ ski -p -c env localhost

    +-----+-----------+--------+----------------+------+--------------------------------------------------------+
    |                                          ski -p -c env localhost                                          |
    +-----+-----------+--------+----------------+------+--------------------------------------------------------+
    | NR. | ID        | TYPE   | CONNECTION     | NAME | OUTPUT                                                 |
    +-----+-----------+--------+----------------+------+--------------------------------------------------------+
    |  1. | localhost | server | root@localhost | Host | SSH_CONNECTION=127.0.0.1 49154 127.0.0.1 22            |
    |     |           |        |                |      | USER=root                                              |
    |     |           |        |                |      | PWD=/root                                              |
    |     |           |        |                |      | HOME=/root                                             |
    |     |           |        |                |      | SSH_CLIENT=127.0.0.1 49154 22                          |
    |     |           |        |                |      | MAIL=/var/mail/root                                    |
    |     |           |        |                |      | SHELL=/bin/bash                                        |
    |     |           |        |                |      | SHLVL=1                                                |
    |     |           |        |                |      | LOGNAME=root                                           |
    |     |           |        |                |      | PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin |
    |     |           |        |                |      | _=/usr/bin/env                                         |
    +-----+-----------+--------+----------------+------+--------------------------------------------------------+

### Templates

Execute a shell or SQL command or script and convert the output based on a [TextFSM][textfsm] template.

    $ ski -s vparams.sql -t vparams db

The SQL script could look like this:

```sql
SET PAGESIZE 0
SET NEWPAGE 0
SET SPACE 0
SET LINESIZE 18000
SET WRAP OFF
SET FEEDBACK OFF
SET ECHO OFF
SET VERIFY OFF
SET HEADING OFF
SET TAB OFF
SET COLSEP ' , '

SELECT NUM, NAME, VALUE FROM V$PARAMETER WHERE NUM IN (526, 530);
```

The template file could look like this:

    $ cat $ORBIT_HOME/templates/vparams.textfsm

    Value Num (\d+)
    Value Name (\S*)
    Value Value (\S*)

    Start
      ^ *${Num}[ |,]*${Name}[ |,]*${Value} -> Record

### Jobs

Bundle command-line arguments to a job to save the report output.

    $ ski -j vparams

The job file could look like this:

    $ cat $ORBIT_HOME/jobs/vparams.skijob

    -s vparam.sql -t vparam db

The report result could look like this:

    $ cat $ORBIT_HOME/reports/vparams/1531410936.skirep

    1531410936
    [["Num", "int"], ["Name", "string"], ["Value", "string"]]
    ["db","Operativ DB",true,["526", "optimizer_adaptive_plans", "FALSE"]]
    ["db","Operativ DB",true,["530", "optimizer_adaptive_statistics", "FALSE"]]

## Development

Clone the repo:

    $ git clone https://github.com/appplant/ski.git && cd ski/

And then execute:

    $ rake compile

To compile the sources locally for the host machine only:

    $ MRUBY_CLI_LOCAL=1 rake compile

You'll be able to find the binaries in the following directories:

- Linux (64-bit Musl): `mruby/build/x86_64-alpine-linux-musl/bin/ski`
- Linux (64-bit GNU): `mruby/build/x86_64-pc-linux-gnu/bin/ski`
- Linux (64-bit, for old distros): `mruby/build/x86_64-pc-linux-gnu-glibc-2.12/bin/ski`
- OS X (64-bit): `mruby/build/x86_64-apple-darwin15/bin/ski`
- Windows (64-bit): `mruby/build/x86_64-w64-mingw32/bin/ski`
- Host: `mruby/build/host2/bin/ski`

For the complete list of build tasks:

    $ rake -T

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/appplant/ski.

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

The code is available as open source under the terms of the [Apache 2.0 License][license].

Made with :yum: in Leipzig

© 2018 [appPlant GmbH][appplant]

[releases]: https://github.com/appplant/ski/releases
[textfsm]: https://github.com/google/textfsm/wiki/TextFSM
[license]: http://opensource.org/licenses/Apache-2.0
[appplant]: www.appplant.de
