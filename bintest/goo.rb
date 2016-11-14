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
require "test/unit"
require 'os'

BIN_PATH = ARGV.fetch(0)



puts File.join(File.dirname(__FILE__), "testtools")
YAMLENV =  Hash["IPS_ORBIT_FILE" => "/go/bintest/test.json"]
PATH = Hash["PATH" => File.join(File.dirname(__FILE__), "testtools")]

class TestGoo < Test::Unit::TestCase

  def test_not_yet_supportet
    output, error, status = Open3.capture3(PATH,BIN_PATH,"id3","command")
    assert_false status.success?, "This type of connection will be implemted at a later state"
    assert_include output, "This Type of Connection is not yet supported"
  end

  def test_not_supportet
    output, error, status = Open3.capture3(PATH,BIN_PATH,"id5","command")
    assert_false status.success?, "Web won't be supported ever"
    assert_include output, "This Type of Connection is not supported"
  end

  def test_commands
    output, error, status = Open3.capture3(PATH,BIN_PATH,"id1","command1","command2","command3")
    assert_true status.success?, "This test should succeed"
    assert_include output, "command2"
  end

  def test_failure
    assert_false(false)
  end

end
=begin
assert('checking json file existence') do
  output, error, status = Open3.capture3(MALFORMEDYAMLENV,BIN_PATH,"id1")

  assert_false status.success?, "Process did not exit cleanly"
  assert_include error, "cannot read from"
end

assert('checking json file correctness') do
  output, error, status = Open3.capture3(MALFORMEDYAMLFILE,BIN_PATH,"id1")

  assert_false status.success?, "Process did not exit cleanly"
  assert_include error, "invalid json"
end

assert('case: missing argument: -t, but no further arguments') do
  output, error, status = Open3.capture3(YAMLENV, BIN_PATH, "-t", "id3")

  assert_false status.success?, "Process did not exit cleanly"
  assert_include error, "Object not found"
end

assert('case: id not found') do
  output, error, status = Open3.capture3(YAMLENV,BIN_PATH,"id404")

  assert_false status.success?, "Process did not exit cleanly"
  assert_include error, "Object not found"
end
assert('case: id found, success') do
  output, error, status = Open3.capture3(YAMLENV,BIN_PATH,"id3")

  assert_true status.success?, "Process did not exit cleanly"
  assert_include output, "user1@url_url1.bla.blergh.de"
end

assert('case: success -t url') do
  output, error, status = Open3.capture3(YAMLENV,BIN_PATH,"-t","url","id4")

  assert_true status.success?, "Process did not exit cleanly"
  assert_include output, "url_url2.bla.blergh.de"
end

assert('case: success -t jdbc') do
  output, error, status = Open3.capture3(YAMLENV,BIN_PATH,"-t","jdbc","id4")

  assert_true status.success?, "Process did not exit cleanly"
  assert_include output, "host2.bla:777:horst2"
end

assert('case: success -t tns') do
  output, error, status = Open3.capture3(YAMLENV,BIN_PATH,"-t","tns","id4")

  assert_true status.success?, "Process did not exit cleanly"
  assert_include output, "(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=TCP)(HOST=host2.bla)(PORT=777)))(CONNECT_DATA=(SID=horst2)))"
end

assert('case: found id, wrong type') do
  output, error, status = Open3.capture3(YAMLENV,BIN_PATH,"id5")

  assert_false status.success?, "Process did not exit cleanly"
  assert_include error, "Wrong type"
end
=end


