## Release Notes: _ski_





### 0.9.2 (not yet released)

JOBS!
# Now supports jobs

Provide jobfile by using 
    ```
    ski -j job.json
    ```
All other flags will be ignored when the -j flag is provided.
The jobfile can be provided as a relative path or as an absolute path.
When provided as a relative path, ski starts looking for it in the folder ORBIT_HOME/jobs
Jobfiles have to be in the following form:
    ```
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
    ```
When running in jobmode, ski writes the output at ORBIT_HOME/jobs_output/$JOBNAME$/$TIMESTAMP$ in the following form:

    ```
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
    ```

# Formatter
Ski now uses Interfaces and a FormatterFactory to dynamically create the right formatter for a job.

# Colors
Ski now colorizes occuring errors in ugly-mode and the whole row of an errored planet in prettymod.

## Further Changes:

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



### 0.9.1 (15.02.2017)

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


### 0.9.0 (12.12.2016)

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
