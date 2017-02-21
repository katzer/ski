## Release Notes

### 0.9.2 (not yet released)

1. 64-bit binary for Linux/BusyBox.

  - Linux (64-bit BusyBox): `build/x86_64-pc-linux-busybox/bin/ski`

2. Compile Linux/GNU binaries with GNU libs (glibc)

# 0.9.1
# Whats new
* Renamed the Project from goo to ski
* Extended functionality, so SQL commands and scripts can be executed on database-type planets
* Supports formatting output via a template into JSON format via the -t flag.

Without formatting:
```
    $ ski -s="showver.sh" app
    willywonka version check 2.1
    -----------------------
    
    willywonka blueberrychocolate factory London OompaLoompa_Headquarters
    -[binaries]---------------------------------------------------------------------
      gateway         4.6.1.1 (build time Oct 21 2015 10:35:11)
      telhandlerkm    4.6.0.2 (build time Oct 13 2016 12:00:31)
      ...
```

With formatting:
```
    $ ski -s="showver.sh" -t="perlver_template" app
    [
    ["willywonka_version", "Section", "Suse", "UnixVersion", "UnixPatch", "Key", "Value", "Key2", "Value2", "Os", "OracleDb"],
    ["2.1", "", "", "", "", "", "", "", "", "", ""],
    ["", "binaries", "", "", "", "", "", "", "", "", ""],
    ...
    ]
```

* Supports formatting output into a prettyprinted ascii-table via the -p flag.

```
    $ ski -s="showver.sh" -t="perlver_template" -p app
    | willywonka_version | Section  | gateway                                   |
     -----------------------------------------------------------------------------
    | 2.1                | binaries | 4.6.1.1 (build time Oct 21 2015 10:35:11) |
```
* Supports loading the users bash profile fur bash remote bash executions.
* Creates a logfile and writes down whats happening.
```
    logfile.log:
    [34mINFO[0m[2017-02-15 15:04:56] Started with args: [./ski -s=shover.sh app]
 
    [31mFATA[0m[2017-02-15 15:04:56] open /root/code/bintest/testFolder/scripts/shover.sh: no such file or directory
    Additional info: called from uploadFile. Keypath: /.ssh/orbit.key
 
    [34mINFO[0m[2017-02-15 15:05:17] Started with args: [./ski -s=showver.sh app]
```

* Looks for ssh-keyfile called orbit.key at ../config/ssh/ if no ORBIT_KEY enviromental variable is provided.
* Scripts to be executed need to be saved in ../scripts/
* Templates to be used need to be saved in ../templates/



#0.9.0
# Whats new

* Added multi-planet-execution-support
* Added option -s="pathToScript" to specify a bash script to be uploaded and executed on a planet
* Added option -p for prettyprinting the output in a neat, readable manner
* Changed the ssh authentification handling in a way that goo looks for a keyfile at $ORBIT_KEY rather than relying on the keyfile being managed by the ssh-manager




# 0.0.1
# Whats new
First Version
* bash command execution per ssh on given planet
