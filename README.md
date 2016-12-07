# goo [![Build Status](https://travis-ci.org/appPlant/goo.svg?branch=master)](https://travis-ci.org/appPlant/goo) [![Build status](https://ci.appveyor.com/api/projects/status/f5imsl77fmg2omba/branch/master?svg=true)](https://ci.appveyor.com/project/katzer/goo/branch/master) [![Code Climate](https://codeclimate.com/github/appPlant/goo/badges/gpa.svg)](https://codeclimate.com/github/appPlant/goo)

Execute commands or collect informations on multiple servers in parallel.

    $ goo -h
    usage: usage: goo [options...] <planet> [<further planets>]... -c="<command>"
    Options:
    -s="<path/to/script>", --script="<path/to/script>"  Execute script and return result
    -p, --pretty                                            Pretty print output as a table
    -t, --type                                              Show type of planet
    -h, --help                                              This help text
    -v, --version                                           Show version number

## Prerequisites
You'll need the following installed and in your `PATH`:
- [ff][ff]

## Installation

Download the latest version from the [release page][releases] and add the executable to your `PATH`.

## Development

Clone the repo:
    
    $ git clone https://github.com/appPlant/goo.git && cd goo/

And then execute:

```bash
$ docker-compose run compile # https://docs.docker.com/engine/installation
```

You'll be able to find the binaries in the following directories:

- Linux (64-bit): `build/x86_64-pc-linux-gnu/bin/goo`
- Linux (32-bit): `build/i686-pc-linux-gnu/bin/goo`
- OS X (64-bit): `build/x86_64-apple-darwin14/bin/goo`
- OS X (32-bit): `build/i386-apple-darwin14/bin/goo`
- Windows (64-bit): `build/x86_64-w64-mingw32/bin/goo`
- Windows (32-bit): `build/i686-w64-mingw32/bin/goo`

## Basic Usage

Get the connection by type:

    $ export ORBIT_FILE=/path/to/orbit.json

    $ goo app-package-1 app-package-2 "hostname"
    $ hostname-1
    $ hostname-2

## Advanced features

Execute a script:

    $ goo -s app-package-1 app-package-2 scripts/hostname.sh
    $ hostname-1
    $ hostname-2

Pretty print output:

    $ goo -p app-package-1 app-package-1 "hostname"
    
      NR   PLANET          hostname            
      ===============================
       0   app-package-1   hostname-1         
       1   app-package-2   hostname-1

## Releases

    $ docker-compose run release

Affer this command finishes, you'll see the /releases for each target in the releases directory.

## Tests

To run all integration tests:

    $ docker-compose run bintest

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
