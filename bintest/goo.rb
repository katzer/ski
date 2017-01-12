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

BIN  = ARGV.fetch(0).freeze
PATH = { 'PATH' => "#{File.expand_path('tools', __dir__)}:#{ENV['PATH']}"  }

# TODO new tests

class TestGoo < Test::Unit::TestCase
  def test_server
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"','-d', 'app')


    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '123'
  end

  def test_web
    _, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"', 'web')

    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, 'not supported'
  end

  def test_not_authorized_host
    _, status = Open3.capture2(PATH, BIN, '-c="echo 123"', 'unauthorized')

    assert_false status.success?, 'Process did exit cleanly'
  end

  def test_offline_host
    _, status = Open3.capture2(PATH, BIN, '-c="echo 123"', 'offline')

    assert_false status.success?, 'Process did exit cleanly'
  end

  def test_help
    output, status = Open3.capture2(PATH, BIN, '-h')

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, 'usage: goo'
  end

  def test_version
    output, status = Open3.capture2(PATH, BIN, '-v')

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '0.9'
  end

  def test_empty_return
    output, status = Open3.capture2(PATH, BIN, '-c="ls -a"', 'app')

    assert_equal output, "", 'return was not empty'
  end
end

