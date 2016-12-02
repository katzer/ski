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


require 'open3'
require 'test/unit'

BIN_PATH = ARGV.fetch(0).freeze

# TODO new tests

class TestGoo < Test::Unit::TestCase
  def test_server
    output, status = Open3.capture2(BIN_PATH, 'app', 'whoami')

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, 'root'
  end

  def test_web
    _, error, status = Open3.capture3(BIN_PATH, 'web', 'whoami')

    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, 'not supported'
  end

  def test_not_authorized_host
    _, status = Open3.capture2(BIN_PATH, 'unauthorized', 'whoami')

    assert_false status.success?, 'Process did exit cleanly'
  end

  def test_offline_host
    _, status = Open3.capture2(BIN_PATH, 'offline', 'whoami')

    assert_false status.success?, 'Process did exit cleanly'
  end
end

