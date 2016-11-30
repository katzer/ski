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

install_glibc() {
    apk add --no-cache wget ca-certificates
    wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub
    wget -q https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.23-r3/glibc-2.23-r3.apk
    apk add --no-cache glibc-2.23-r3.apk
}

install_deps() {
    # MacPorts
    if which port >/dev/null; then
        port selfupdate
        port install go ruby23 git
    # Homebrew
    elif which brew >/dev/null; then
        brew update
        brew install go ruby git
    # Alpine-Linux
    elif which apk >/dev/null; then
        apk update
        apk add --no-cache ruby openssh git
        install_glibc
    fi
}

install_pkgs() {
    gem install rake os test-unit --no-ri --no-rdoc
    go get gopkg.in/hypersleep/easyssh.v0
}

setup_ssh_server() {
    /usr/bin/ssh-keygen -A
    /usr/sbin/sshd
    mkdir $HOME/.ssh
    ssh-keygen -q -f $HOME/.ssh/orbit_rsa -N ""
    cp $HOME/.ssh/orbit_rsa.pub $HOME/.ssh/authorized_keys
}

install_deps
install_pkgs
setup_ssh_server
