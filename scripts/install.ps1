#!/bin/bash

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

function Install-Deps() {
    choco install ruby golang openssh git -y
    $env:PATH=[System.Environment]::GetEnvironmentVariable("Path","Machine")
}

function Update-Certs() {
    iwr https://rubygems.org/downloads/rubygems-update-2.6.8.gem -O C:\rubygems-update-2.6.8.gem
    gem install --local C:\rubygems-update-2.6.8.gem --no-ri --no-rdoc
    update_rubygems --no-ri --no-rdoc
    rm C:\rubygems-update-2.6.8.gem
}

function Install-Pkgs() {
    Update-Certs
    gem install rake os test-unit --no-ri --no-rdoc
    if (-not (Test-Path env:GOPATH)) { $env:GOPATH = "$HOME\goms" }
    go get gopkg.in/appPlant/easyssh.v0
    go get github.com/sirupsen/logrus
    go get github.com/rifflock/lfshook
    go get github.com/lestrrat/go-file-rotatelogs
}

function Setup-Sshd() {
    cd 'C:\Program Files\OpenSSH-Win64'
    .\install-sshd.ps1
    .\install-sshlsa.ps1
    ssh-keygen -A
    ntrights.exe -u "NT SERVICE\SSHD" +r SeAssignPrimaryTokenPrivilege
    mkdir -f $HOME\.ssh
    Remove-Item $HOME\.ssh\orbit.key -ErrorAction Ignore
    cd 'C:\Program Files\Git\usr\bin'
    .\ssh-keygen -q -f $HOME\.ssh\orbit.key -N "''"
    cp $HOME\.ssh\orbit.key.pub $HOME\.ssh\authorized_keys
    Start-Service sshd
    .\ssh-keyscan -t ecdsa-sha2-nistp256 localhost 2> $NULL > $HOME\.ssh\known_hosts
    # New-NetFirewallRule -Protocol TCP -LocalPort 22 -Direction Inbound -Action Allow -DisplayName SSH
    # netsh advfirewall firewall add rule name='SSH Port' dir=in action=allow protocol=TCP localport=22
}

Install-Deps
Install-Pkgs
Setup-Sshd
