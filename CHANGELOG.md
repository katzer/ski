# Release Notes: _ski_

Execute commands or collect informations on multiple servers in parallel.

## 1.5.1

Released at: 18.03.2020

1. Singularized folder names

2. Fixed potential memory leaks.

3. Compiled with `MRB_WITHOUT_FLOAT`

4. Compiled binary for OSX build with MacOSX10.15 SDK

5. Upgraded to mruby 2.1.0

[Full Changelog](https://github.com/appplant/ski/compare/1.5.0...1.5.1)

## 1.5.0

Released at: 13.08.2019

<details><summary>Releasenotes</summary>
<p>

1. Added support for `ECDSA` for both key exchange and host key algorithms

2. Compiled binary for OSX build with MacOSX10.13 SDK (Darwin17)

3. Upgraded to mruby 2.0.1

</p>

[Full Changelog](https://github.com/appplant/ski/compare/1.4.7...1.5.0)
</details>

## 1.4.7

Released at: 02.01.2019

<details><summary>Releasenotes</summary>
<p>

1. Dropped compatibility with orbit v1.4.6 due to breaking changes in _fifa_.

2. Removed LVAR section for non test builds.

3. Upgraded to mruby 2.0.0

</p>

[Full Changelog](https://github.com/appplant/ski/compare/1.4.6...1.4.7)
</details>

## 1.4.6

Released at: 16.08.2018

<details><summary>Releasenotes</summary>
<p>

Tool has been fully reworked!

    $ ski -h

    usage: ski [options...] matchers...
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

#### Templates

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

#### Jobs

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

</p>

[Full Changelog](https://github.com/appplant/ski/compare/1.4.4...1.4.6)
</details>

## 1.4.4

Released at: 30.11.2017

<details><summary>Releasenotes</summary>
<p>

Provide jobfile by using 

    $ ski -j job.json

All other flags will be ignored when the -j flag is provided.
The jobfile can be provided as a relative path or as an absolute path.
When provided as a relative path, ski starts looking for it in the folder ORBIT_HOME/jobs
Jobfiles have to be in the following form:

    {
        "debug":true,
        "help":false,
        "load":false,
        "pretty":false,
        "version":false,
        "save_report":false,
        "command":"ls -a",
        "scriptName":"",
        "template":"",
        "planets":[
            "app",
            "app"
        ],
        "LogFile":""
    }

When running in jobmode, ski writes the output at ORBIT_HOME/jobs_output/$JOBNAME$/$TIMESTAMP$ in the following form:

    {
        "meta": {
            "debug": true,
            "help": false,
            "load": false,
            "pretty": true,
            "version": false,
            "save_report": false,
            "command": "ls -a",
            "scriptName": "",
            "template": "",
            "planets": [
                "app",
                "app"
            ],
            "log_file": ""
        },
        "planets": [
            {
                "id": "app",
                "output": ".\n..\n.bash_profile\n.bashrc\n.gem\n.gitconfig\n.profile\n.ssh\ncode\nprofiles\nsql\n",
            },
            {
                "id": "app",
                "output": ".\n..\n.bash_profile\n.bashrc\n.gem\n.gitconfig\n.profile\n.ssh\ncode\nprofiles\nsql\n",
            }
        ]
    }

### Formatter
Ski now uses Interfaces and a FormatterFactory to dynamically create the right formatter for a job.

### Colors
Ski now colorizes occuring errors in ugly-mode and the whole row of an errored planet in prettymod.

### Further Changes:

1. 64-bit binary for Linux/BusyBox.

2. Compile Linux/GNU binaries with GNU libs (glibc).

3. Strip binaries to shrink their size.

4. errors within planets don't cause an abort, but cause an abort for that planet alone and writes the error in the planets output

5. ski now validates fifas output

6. fifa is now being called a single time, retrievingevery possible information about every planet

7. fifa mockup updated to return multiple values

8. Output by automatic addition of a semicolon to a db command changed to log warning

9. stripped metadata from db output

10. now checks if a given script ends in a supported extension(supported extensions are .sh and .sql)

11. removed structuredOutputList from ski.go. Now everything is being handled with a planetlist with every planet containing its own StructuredOutput

12. structured output now contains a variable indicating its planets position

13. changed most functions requiring a structuredOutput

14. prettyprint now uses package ascii table rather than the self implemented version

15. added field "address" to prettytable and pretty

16. checks templates existance

17. Expanded prettytables functions so it's able to display multiple planets in a single table, indicating missing values with a "-"

18. expanded prettytable with metadata

19. optionparsing moved from optparser to ski.go

20. Files uploaded to the ssh target are now modified with the planets position, so conflicts with running a script multiple times on one planet are avoided

21. ski is now able to parse job configurations from a file

22. inserted codebeat disables for planet and opts

23. removed unnecessary comments

24. removed deprecated functions

</p>

[Full Changelog](https://github.com/appplant/ski/compare/0.9.1...1.4.4)
</details>

## 0.9.1

Released at: 15.02.2017

<details><summary>Releasenotes</summary>
<p>

1. Renamed the tool to ski (<b>S</b>ascha <b>K</b>nows <b>I</b>t).

2. Extended functionality, so SQL commands and scripts can be executed on databases.

   ```
   $ ski -c="SELECT * FROM DUAL;" ora-db

   D
   -
   X
   ```

3. Added -l to load the user profile:

   ```
   $ ski -l -c="padm d" app-package-1
   ```

4. Look for orbit.key at $ORBIT_HOME/config/ssh/ if no $ORBIT_KEY env is provided.

3. Added -t to convert the output into JSON by using a [TextFSM](https://github.com/google/textfsm/wiki/TextFSM) template:

   ```
   $ ski -s="showver.sh" -t="showver" app

   [
   ["showver_version", "Section", "Suse", "UnixVersion", "UnixPatch", "Key", "Value", "Key2", "Value2", "Os", "OracleDb"],
   ["2.1", "", "", "", "", "", "", "", "", "", ""],
   ["", "binaries", "", "", "", "", "", "", "", "", ""],
   ...
   ]
   ```

4. Supports formatting template output into a prettyprinted ascii-table:

   ```
   $ ski -s="showver.sh" -t="showver" -p app

    | showver_version | Section  | gateway                                   |
    |-----------------|----------|-------------------------------------------|
    | 2.1             | binaries | 4.6.1.1 (build time Oct 21 2015 10:35:11) |
   ```

5. Logging to see whats happening:

   ```
   $ cat $ORBIT_HOME/log/logfile.log

   [34mINFO[0m[2017-02-15 15:04:56] Started with args: [./ski -s=shover.sh app]
   [31mFATA[0m[2017-02-15 15:04:56] open /root/code/bintest/testFolder/scripts/shover.sh: no such file or directory
   Additional info: called from uploadFile. Keypath: /.ssh/orbit.key
   [34mINFO[0m[2017-02-15 15:05:17] Started with args: [./ski -s=showver.sh app]
   ```

</p>

[Full Changelog](https://github.com/appplant/ski/compare/0.9.0...0.9.1)
</details>

## 0.9.0

Released at: 12.12.2016

<details><summary>Releasenotes</summary>
<p>

1. Execute command/script/sql on multiple planets at same time:

   ```
   $ ski -c="hostname" app-package-1 app-package-2

   hostname-1
   hostname-2
   ```

2. Added -s to specify a bash script to be executed:

   ```
   $ ski -s="scripts/hostname.sh" app-package-1 app-package-2

   hostname-1
   hostname-2
   ```

3. Added -p for prettyprinting the output in a neat, readable manner.

   ```
   $ ski -p -c="hostname" app-package-1 app-package-2

     NR   PLANET          hostname
     ===============================
      0   app-package-1   hostname-1
      1   app-package-2   hostname-1
   ```

4. Changed the ssh authentification handling in a way that goo looks for a keyfile at $ORBIT_KEY rather than relying on the keyfile being managed by the ssh-manager.

</p>

[Full Changelog](https://github.com/appplant/ski/compare/1405d1037363ff32e8c2bf28664efd9896859631...0.9.0)
</details>
