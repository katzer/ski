# goo [![Build Status](https://travis-ci.org/appPlant/goo.svg?branch=master)](https://travis-ci.org/appPlant/goo) [![Build status](https://ci.appveyor.com/api/projects/status/f5imsl77fmg2omba/branch/master?svg=true)](https://ci.appveyor.com/project/katzer/goo/branch/master) [![Code Climate](https://codeclimate.com/github/appPlant/goo/badges/gpa.svg)](https://codeclimate.com/github/appPlant/goo)

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
    
    $ git clone https://github.com/appPlant/goo.git && cd goo/

And then execute:

```bash
$ scripts/compile # https://docs.docker.com/engine/installation
```

You'll be able to find the binaries in the following directories:

- Linux (64-bit, for old distros): `build/x86_64-pc-linux-gnu-glibc-2.12/bin/ski`
- Linux (32-bit, for old distros): `build/i686-pc-linux-gnu-glibc-2.12/bin/ski`
- Linux (64-bit GNU): `build/x86_64-pc-linux-gnu-glibc-2.14/bin/ski`
- Linux (32-bit GNU): `build/i686-pc-linux-gnu-glibc-2.14/bin/ski`
- Linux (64-bit BusyBox): `build/x86_64-pc-linux-busybox-musl/bin/ski`
- OS X (64-bit): `build/x86_64-apple-darwin14/bin/ski`
- OS X (32-bit): `build/i386-apple-darwin14/bin/ski`
- Windows (64-bit): `build/x86_64-w64-mingw32/bin/ski`
- Windows (32-bit): `build/i686-w64-mingw32/bin/ski`

## Basic Usage

Get the connection by type:

    $ export ORBIT_FILE=/path/to/orbit.json

    $ ski -c="hostname" app-package-1 app-package-2
    $ hostname-1
    $ hostname-2

## Advanced features

Execute a script:

    $ ski -s="scripts/hostname.sh" app-package-1 app-package-2
    $ hostname-1
    $ hostname-2

Pretty print output:

    $ ski -p -c="hostname" app-package-1 app-package-2
    
      NR   PLANET          hostname
      ===============================
       0   app-package-1   hostname-1
       1   app-package-2   hostname-1

## Releases

    $ scripts/release

Affer this command finishes, you'll see the /releases for each target in the releases directory.

## Tests

To run all integration tests:

    $ scripts/bintest

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/appPlant/goo.

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
[releases]: https://github.com/appPlant/goo/releases
[docker]: https://docs.docker.com/engine/installation
[license]: http://opensource.org/licenses/Apache-2.0
[appplant]: www.appplant.de
