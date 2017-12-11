# ski [![Build Status](https://travis-ci.org/appPlant/ski.svg?branch=master)](https://travis-ci.org/appPlant/ski) [![Build status](https://ci.appveyor.com/api/projects/status/f5imsl77fmg2omba/branch/master?svg=true)](https://ci.appveyor.com/project/katzer/goo/branch/master) [![codebeat badge](https://codebeat.co/badges/b0a926f1-d7bf-4ee1-9bc8-4cb1e087d347)](https://codebeat.co/projects/github-com-appplant-ski)

Execute commands or collect informations on multiple servers in parallel.

    $ ski -h
    usage: ski [options...] -c="<command>" <planets>... 
    Options:
    -s="<scriptname>"   	Execute script and return result
    -c="<command>"  	    Execute script and return result
    -t=<"templatename>" 	Templatefile to be applied 
    -p    			        Pretty print output as a table
    -l    			        Load bash profiles on Server
    -t    			        Show type of planet
    -h    			        Display this help text
    -v    			        Show version number
    -d			            Show extended debug informations


## Prerequisites
Create an enviroment variable called `ORBIT_HOME` and set it to the absolute path of the ski folder-structure. 

Example: You save the release at `/home/youruser/workspace/ski`. Your `ORBIT_HOME` should be `/home/youruser/workspace/ski`as well

Either create an enviroment variable called `ORBIT_KEY` containing an absolute path to the ssh private key that should be used for executing commands on the planets or save said key at `ski/config/ssh/`.

You'll need the following installed and in your `PATH`:
- [fifa][ff]

## Installation

Download the latest version from the [release page][releases] and add the executable to your `PATH`.

## Development

Clone the repo:
    
    $ git clone https://github.com/appPlant/ski.git && cd ski/

And then execute:

```bash
$ scripts/compile # https://docs.docker.com/engine/installation
```

You'll be able to find the binaries in the following directories:

- Linux (64-bit, for old distros): `build/x86_64-pc-linux-gnu-glibc-2.12/bin/ski`
- Linux (32-bit, for old distros): `build/i686-pc-linux-gnu-glibc-2.12/bin/ski`
- Linux (64-bit GNU): `build/x86_64-pc-linux-gnu-glibc-2.14/bin/ski`
- Linux (32-bit GNU): `build/i686-pc-linux-gnu-glibc-2.14/bin/ski`
- Linux (64-bit Musl): `build/x86_64-alpine-linux-musl/bin/ski`
- OS X (64-bit): `build/x86_64-apple-darwin15/bin/ski`
- OS X (32-bit): `build/i386-apple-darwin15/bin/ski`
- Windows (64-bit): `build/x86_64-w64-mingw32/bin/ski`
- Windows (32-bit): `build/i686-w64-mingw32/bin/ski`

## Basic Usage

Execute commands or collect informations on multiple servers in parallel.

    $ ski -h
    usage: ski [options...] -c="<command>" <planets>... 
    Options:
    -s="<scriptname>"   	        Execute script and return result
    -c="<command>"  	        Execute script and return result
    -t=<"templatename>" 	        Templatefile to be applied 
    -p    			        Pretty print output as a table
    -l    			        Load bash profiles on Server
    -h    			        Display this help text
    -v    			        Show version number
    -d			            Show extended debug informations

#### Command execution - Linux server
Use ski to execute commands on linux server planets:
```
$ ski -c="echo hi" app-package-1
hi
```

#### Command execution - Database
Use ski to execute commands on database planets:
```
$ ski -c="SELECT * FROM DUAL;" ora-db
D
-
X
```
#### Script execution - Linux server
Use ski to execute bash scripts on linux server planets:
```
$ ski -s="echo-script.sh" app-package-1
hi
```

#### Script execution - Database
Use ski to execute sql srcipts on database planets:
```
$ ski -s="select.sql" ora-db
D
-
X
```

#### Prettyprinting script and command execution
Set the pretty flag to display the planet output from any script or command execution in a neat, readable manner:
```
$ ski -c="echo hi" -p app
+-----+-----+---------------+-------------------+--------+---------------------------------------+
| NR  | ID  |     NAME      |      ADDRESS      |  TYPE  |                OUTPUT                 |
+-----+-----+---------------+-------------------+--------+---------------------------------------+
|   0 | app | App-Package 1 | user@localhost    | server |                  hi                   |
+-----+-----+---------------+-------------------+--------+---------------------------------------+
```

#### Output conversion to JSON - template
Provide a [TextFSM](https://github.com/google/textfsm/wiki/TextFSM) template to convert the output of a planet to json:
```
$ ski -s="showver.sh" -t="showver" app

[
["showver_version", "Section", "Suse", "UnixVersion", "UnixPatch", "Key", "Value", "Key2", "Value2", "Os", "OracleDb"],
["2.1", "", "", "", "", "", "", "", "", "", ""],
["", "binaries", "", "", "", "", "", "", "", "", ""],
...
]
```

#### Prettyprinting output converted to JSON
Set the pretty flag to display the planet output from any [TextFSM](https://github.com/google/textfsm/wiki/TextFSM) template converted execution in a neat, readable manner:
```
$ ski -s="showver.sh" -t="showver" -p app

| showver_version | Section  | gateway                                   |
|-----------------|----------|-------------------------------------------|
| 2.1             | binaries | 4.6.1.1 (build time Oct 21 2015 10:35:11) |
```
#### Jobs
Provide a JSON jobfile to persistently save a configuration to run ski with. Any flag not set in this jobfile will be considered not provided. This way complicated ski execution settings can be stored and called with providing but one parameter to ski.
The output from this job will be stored in an automatically created folder called `job_output`

Usage: 
```
$ ski -j="showver.json" app
```

```
showver.json:

{
  "pretty"     : true,
  "scriptName" : "showver.sh",
  "template"   : "showver",
  "planets"    : [ "app-1", "app-2", "app-3", "app-5" ]
}
```

## Releases

    $ scripts/release

Affer this command finishes, you'll see the /releases for each target in the releases directory.

## Tests

To run all integration tests:

    $ scripts/bintest

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/appPlant/ski.

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


## License

The code is available as open source under the terms of the [Apache 2.0 License][license].

Made with :yum: from Leipzig

© 2016 [appPlant GmbH][appplant]

[ff]: https://github.com/appPlant/ff/releases
[releases]: https://github.com/appPlant/ski/releases
[docker]: https://docs.docker.com/engine/installation
[license]: http://opensource.org/licenses/Apache-2.0
[appplant]: www.appplant.de
