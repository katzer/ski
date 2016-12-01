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

$env:GOROOT = "C:\go"
$env:GOPATH = "C:\go\pkgs"

function Install-Deps() {
    choco install ruby golang openssh git -y
    $env:PATH = [System.Environment]::GetEnvironmentVariable("Path","Machine")
}

function Update-Certs() {
    iwr https://rubygems.org/downloads/rubygems-update-2.6.8.gem -O C:\rubygems-update-2.6.8.gem
    gem install --local C:\rubygems-update-2.6.8.gem --no-ri --no-rdoc
    update_rubygems --no-ri --no-rdoc
    rm C:\rubygems-update-2.6.8.gem
}

function Install-Pkgs() {
    gem install rake os test-unit --no-ri --no-rdoc
    go get gopkg.in/hypersleep/easyssh.v0
}

Install-Deps
Update-Certs
Install-Pkgs