## Release Notes: _ski_

### 0.9.2 (not yet released)

1. 64-bit binary for Linux/BusyBox.

2. Compile Linux/GNU binaries with GNU libs (glibc).

3. Strip binaries to shrink their size.


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
