#!/bin/sh

export ORBIT_HOME=/code/bintest/testFolder
rm ./goo
cp /code/build/x86_64-pc-linux-gnu/bin/goo ./goo
./goo -s="/code/bintest/tools/showver.sh" -tn="showverTemp" -pyp="/code/bintest/testFolder/pythonScripts" app
