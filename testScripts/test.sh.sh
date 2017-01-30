#!/bin/sh

<<<<<<< HEAD
export ORBIT_HOME=/code/bintest/testFolder
rm /code/bintest/testFolder/bin/goo
cp /code/build/x86_64-pc-linux-gnu/bin/goo /code/bintest/testFolder/bin/goo

/code/bintest/testFolder/bin/goo -s="test.sh" app
=======
export ORBIT_HOME=/root/code/bintest/testFolder
rm /root/code/bintest/testFolder/bin/goo
cp /root/code/build/x86_64-pc-linux-gnu/bin/goo /root/code/bintest/testFolder/bin/goo

/root/code/bintest/testFolder/bin/goo -s="test.sh" -d app
>>>>>>> f5eac5922a8f6bd4412d4a126ee646d164e48fb2
