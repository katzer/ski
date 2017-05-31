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

require 'rubygems'
require 'os'
require 'go/build'

Go::Build.new('x86_64-pc-linux-gnu') do
  os :linux
  arch :amd64
  appname :ski
  bintest_if OS.linux? && OS.bits == 64
end

Go::Build.new('i686-pc-linux-gnu') do
  os :linux
  arch :'386'
  appname :ski
  bintest_if OS.linux? && OS.bits == 32
end

Go::Build.new('x86_64-apple-darwin15') do
  os :darwin
  arch :amd64
  appname :ski
  bintest_if OS.mac? && OS.bits == 64
end

Go::Build.new('i386-apple-darwin15') do
  os :darwin
  arch :'386'
  appname :ski
  bintest_if OS.mac? && OS.bits == 32
end

Go::Build.new('x86_64-w64-mingw32') do
  os :windows
  arch :amd64
  appname :"ski.exe"
  bintest_if OS.windows? && OS.bits == 64
end

Go::Build.new('i686-w64-mingw32') do
  os :windows
  arch :'386'
  appname :"ski.exe"
  bintest_if OS.windows? && OS.bits == 32
end
