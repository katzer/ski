#!/bin/sh

#
# Copyright (c) 2013-2016 by appPlant GmbH. All rights reserved.
#
# @APPPLANT_LICENSE_HEADER_START@
#
# This file contains Original Code and/or Modifications of Original Code
# as defined in and that are subject to the Apache License
# Version 2.0 (the 'License'). You may not use this file except in
# compliance with the License. Please obtain a copy of the License at
# http://opensource.org/licenses/Apache-2.0/ and read it before using this
# file.
#
# The Original Code and all software distributed under the License are
# distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
# EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
# INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
# Please see the License for the specific language governing rights and
# limitations under the License.
#
# @APPPLANT_LICENSE_HEADER_END@

init_go() {
    [ -z "$GOROOT" ] && export GOROOT=/usr/local/go
    [ -z "$GOPATH" ] && export GOPATH=/go
    export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
}

init_orbit() {
    export ORBIT_KEY=/.ssh/orbit.key
    export ORBIT_HOME=`pwd`/bintest/testFolder
    export PATH=`pwd`/bintest/tools:$PATH
    chmod -R u+x `pwd`/bintest/tools
}

init_sshd() {
    if [[ `id -u` -ne 0 ]]
        then sudo /usr/sbin/sshd
        else /usr/sbin/sshd
    fi
    eval `ssh-agent -s`
    ssh-add $HOME$ORBIT_KEY
}

init_go
init_orbit
init_sshd
git -C $GOPATH/src/golang.org/x/crypto reset --hard abc5fa7ad02123a41f02bf1391c9760f7586e608
