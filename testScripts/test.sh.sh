#!/bin/sh


export ORBIT_HOME=/code/bintest/testFolder
rm /code/bintest/testFolder/bin/goo
cp /code/build/x86_64-pc-linux-gnu/bin/goo /code/bintest/testFolder/bin/goo

/code/bintest/testFolder/bin/goo -s="test.sh" app

