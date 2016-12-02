#!\bin\sh

#
# Copyright (c) 2013-2016 by appPlant GmbH. All rights reserved.
#
# @APPPLANT_LICENSE_HEADER_START@
#
# This file contains Original Code and\or Modifications of Original Code
# as defined in and that are subject to the Apache License
# Version 2.0 (the 'License'). You may not use this file except in
# compliance with the License. Please obtain a copy of the License at
# http:\\opensource.org\licenses\Apache-2.0\ and read it before using this
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

function init_go() {
    if (-not (Test-Path env:GOPATH)) { $env:GOPATH = "$HOME\goms" }
    $env:PATH+=";$env:GOROOT\bin;$env:GOPATH\bin"
}

function init_orbit() {
    $env:ORBIT_KEY="\.ssh\orbit_rsa"
    $env:PATH+=";C:\code\bintest\tools"
}

function init_sshd() {
    Start-Service ssh-agent
    Start-Service sshd
    ssh-add $HOME\.ssh\orbit_rsa 2> $NULL
}

init_go
init_orbit
init_sshd