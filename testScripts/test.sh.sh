#!/bin/sh


export ORBIT_HOME=/code/bintest/testFolder
rm /root/code/bintest/testFolder/bin/goo
cp /root/code/build/x86_64-pc-linux-gnu/bin/goo /root/code/bintest/testFolder/bin/goo

/root/code/bintest/testFolder/bin/goo -s="test.sh" app
