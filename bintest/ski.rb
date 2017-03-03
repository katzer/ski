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


class TestGoo < Test::Unit::TestCase
  def test_server
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"',
                                           '-d=true', 'app')
    check_error(output, error, 'test_server')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '123'
  end

  def test_web
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"',
                                           '-d=true', 'web')
    check_no_error(output, error, 'test_web')
    assert_true status.success?, 'Process did exit cleanly'
    assert_include error, 'Usage of ski with web servers is not implemented'
  end

  # def test_not_authorized_host
  #   output, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"',
  #                                          '-d=true', 'unauthorized')
  #   check_no_error(output, error, 'test_not_authorized_host')
  #   assert_false status.success?, 'Process did exit cleanly'
  #   assert_include error, 'ssh: unable to authenticate'
  # end

  # def test_offline_host
  #   output, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"',
  #                                          '-d=true', 'offline')
  #   check_no_error(output, error, 'test_offline_host')
  #   assert_false status.success?, 'Process did exit cleanly'
  #   assert_include error, 'no such host'
  # end

  def test_help
    output, error, status = Open3.capture3(PATH, BIN, '-h')
    check_error(output, error, 'test_help')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, 'usage: ski'
  end

  def test_version
    output, error, status = Open3.capture3(PATH, BIN, '-v')
    check_error(output, error, 'test_version')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '0.9'
  end

  def test_empty_return
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo "',
                                           '-d=true', 'app')
    check_error(output, error, 'test_empty_return')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_equal output, "\n", 'return was not empty'
  end

  def test_table_print
    output, error, status = Open3.capture3(PATH, BIN, '-s="showver.sh"',
                                           '-t="perlver_template"',
                                           '-d=true', 'app')
    check_error(output, error, 'test_tablePrint')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, "\n[\"willywonka_version\",", 'return was not right'
  end

  def test_pretty_table_print
    output, error, status = Open3.capture3(PATH, BIN, '-s="showver.sh"',
                                           '-t="perlver_template"', '-p',
                                           '-d=true', 'app')
    check_error(output, error, 'test_pretty_tablePrint')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '| WILLYWONKA VERSION |', 'return was not right'
  end

  def test_script_execution
    output, error, status = Open3.capture3(PATH, BIN, '-s="test.sh"',
                                           '-d=true', 'app')
    check_error(output, error, 'test_script_execution')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_equal output, "bang\n", 'return was not correct'
  end

  def test_no_such_script
    output, error, status = Open3.capture3(PATH, BIN, '-s="nonExistent.sh"',
                                           '-d=true', 'app')
    check_no_error(output, error, 'no_such_script')
    assert_true status.success?, 'Process did exit cleanly'
    assert_equal output, '', 'return was not correct'
    assert_include error, 'no such file or directory', 'error was not correct'
  end

  def test_bad_script
    output, error, status = Open3.capture3(PATH, BIN, '-s="badscript.sh"',
                                           'app')
    check_no_error(output, error, 'bad_script')
    assert_true status.success?, 'Process did exit cleanly'
    assert_include error, 'Process exited with status 127', 'return incorrect'
  end

  def test_bad_command
    output, error, status = Open3.capture3(PATH, BIN, '-c="yabeda baba"',
                                           '-d=true', 'app')
    check_no_error(output, error, 'bad_command')
    assert_true status.success?, 'Process did exit cleanly'
    assert_include error, 'Process exited with status 127', 'return incorrect'
  end

  def test_pretty_print
    output, error, status = Open3.capture3(PATH, BIN, '-c="ls -al"', '-p',
                                           '-d=true', 'app')
    check_error(output, error, 'pretty_print')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '|   0 | app       |', 'return was incorrect'
  end

  def test_multiple_pretty_print
    output, error, status = Open3.capture3(PATH, BIN, '-c="ls -al"', '-p',
                                           '-d=true', 'app', 'app', 'app')
    check_error(output, error, 'pretty_print')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '|   0 | app       |', 'return was incorrect'
    assert_include output, '|   1 | app       |', 'return was incorrect'
    assert_include output, '|   2 | app       |', 'return was incorrect'
  end

  def test_malformed_flag
    output, error, status = Open3.capture3(PATH, BIN, '-c="ls -al"', '-zz',
                                           '-d=true', 'app')
    check_no_error(output, error, 'malformed_flag')
    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, 'Usage of', 'return was not correct'
  end

  def test_not_enough_args
    output, error, status = Open3.capture3(PATH, BIN, '-p', '-d=true', 'app')
    check_error(output, error, 'not_enough_args')
    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, 'usage:', 'return was not correct'
  end

  def test_wrong_flag_order
    output, error, status = Open3.capture3(PATH, BIN, '-c="ls -al"', 'app',
                                           '-d=true', '-p')
    check_no_error(output, error, 'wrong_flag_order')
    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, 'Unknown target', 'error was not correct'
  end

  def test_nonexistent_planet
    output, error, status = Open3.capture3(PATH, BIN, '-c="ls -al"', '-d=true',
                                           'pep')
    check_no_error(output, error, 'nonexistent_planet')
    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, 'Unknown target', 'error was not correct'
  end

  def test_no_template
    output, error, status = Open3.capture3(PATH, BIN, '-s="showver.sh"',
                                           '-t="no_template"', '-d=true',
                                           '-p', 'app')
    check_no_error(output, error, 'no_template')
    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, 'exit status 2', 'wrong error'
  end

  def test_malformed_template
    output, error, status = Open3.capture3(PATH, BIN, '-s="showver.sh"',
                                           '-t="useless_template"', '-d=true',
                                           '-p', 'app')
    check_no_error(output, error, 'malformed_template')
    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, 'exit status 2', 'wrong error'
  end

  def test_copy_failed
    command = '-c="touch test && cp test ./test/test"'
    output, error, status = Open3.capture3(PATH, BIN, command,
                                           '-p', '-d=true', 'app')
    check_no_error(output, error, 'copy_failed')
    assert_true status.success?, 'Process did exit cleanly'
    assert_include error, 'Process exited with status 1', 'wrong error'
  end

#  def test_glibc_version
#    output, = Open3.capture2 "readelf -V #{BIN} | grep GLIBC_2.[1-9][0-9]"
#
#    assert_empty output
#  end
end

def check_error(output, error, test_name)
  return if error.empty?
  puts "test: #{test_name}"
  puts "output: #{output}"
  puts "error: #{error}"
end

def check_no_error(output, error, test_name)
  return unless error.empty?
  puts "test: #{test_name}"
  puts "output: #{output}"
  puts "error: #{error.inspect}"
end
