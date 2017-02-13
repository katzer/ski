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
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"', 'app')

    checkError(output,error,"test_server")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '123'
  end

  def test_web
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"', 'web')

    checkNoError(output,error,"test_web")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include error, 'Usage of goo with web servers is not implemented'
  end

  def test_not_authorized_host
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"', 'unauthorized')

    checkNoError(output,error,"test_not_authorized_host")

    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, 'ssh: unable to authenticate'
  end

  def test_offline_host
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo 123"', 'offline')

    checkNoError(output,error,"test_offline_host")

    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, 'no such host'
  end

  def test_help
    output, error, status = Open3.capture3(PATH, BIN, '-h')

    checkError(output,error,"test_help")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, 'usage: goo'
  end

  def test_version
    output, error, status = Open3.capture3(PATH, BIN, '-v')

    checkError(output,error,"test_version")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, '0.9'
  end

  def test_empty_return
    output, error, status = Open3.capture3(PATH, BIN, '-c="echo "', 'app')

    checkError(output,error,"test_empty_return")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_equal output, "\n", 'return was not empty'
  end


  def test_tablePrint
    toolsPath = File.expand_path('tools', __dir__)
    output, error, status = Open3.capture3(PATH, BIN,"-s=\"showver.sh\"", "-t=\"perlver_template\"", "app")

    checkError(output,error,"test_tablePrint")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, "\n[\"IPST_Version\", \"Section\", \"Suse\", \"UnixVersion\", \"UnixPatch\", \"Key\", \"Value\", \"Key2\", \"Value2\", \"Os\", \"OracleDb\"],\n", 'return was not right'
  end

  def test_pretty_tablePrint
    toolsPath = File.expand_path('tools', __dir__)
    output, error, status = Open3.capture3(PATH, BIN,"-s=\"showver.sh\"", "-t=\"perlver_template\"","-p", "app")

    checkError(output,error,"test_pretty_tablePrint")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, "| IPST_Version | Section  | gateway                                   | telhandlerkm                              |", 'return was not right'
  end



  def test_script_execution
    output, error, status = Open3.capture3(PATH, BIN, "-s=\"test.sh\"", 'app')

    checkError(output,error,"test_script_execution")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_equal output, "bang\n", 'return was not correct'
  end

  def test_no_such_script
    output, error, status = Open3.capture3(PATH, BIN, "-s=\"nonExistent.sh\"", 'app')

    checkNoError(output,error,"no_such_script")

    assert_false status.success?, 'Process did exit cleanly'
    assert_equal output, "", 'return was not correct'
    assert_include error, "no such file or directory", 'error was not correct'
  end

  def test_bad_script
    output, error, status = Open3.capture3(PATH, BIN, "-s=\"badscript.sh\"", 'app')

    checkNoError(output,error,"bad_script")

    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, "Process exited with status 127", 'return was not correct'
  end

  def test_bad_command
    output, error, status = Open3.capture3(PATH, BIN, "-c=\"yabeda baba\"", 'app')

    checkNoError(output,error,"bad_command")

    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, "Process exited with status 127", 'return was not correct'
  end

  def test_pretty_print
    output, error, status = Open3.capture3(PATH, BIN, "-c=\"ls -al\"","-p", 'app')

    checkError(output,error,"pretty_print")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, "0    app                  ", 'return was not correct'
  end

  def test_multiple_pretty_print
    output, error, status = Open3.capture3(PATH, BIN, "-c=\"ls -al\"","-p", 'app','app','app')

    checkError(output,error,"pretty_print")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, "0    app", 'return was not correct'
    assert_include output, "1    app", 'return was not correct'
    assert_include output, "2    app", 'return was not correct'
  end

  def test_malformed_flag
    output, error, status = Open3.capture3(PATH, BIN, "-c=\"ls -al\"","-zz", 'app')

    checkNoError(output,error,"malformed_flag")

    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, "Usage of", 'return was not correct'
  end

  def test_wrong_flag_order
    output, error, status = Open3.capture3(PATH, BIN, "-c=\"ls -al\"", 'app', "-p")

    checkNoError(output,error,"wrong_flag_order")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, "total", 'return was not correct'
    assert_include error, "Unkown Type of target", 'error was not correct'
  end

  def test_nonexistent_planet
    output, error, status = Open3.capture3(PATH, BIN, "-c=\"ls -al\"", 'pep')

    checkNoError(output,error,"nonexistent_planet")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include error, "Unkown Type of target", 'error was not correct'
  end

  def test_not_enough_args
    output, error, status = Open3.capture3(PATH, BIN, "-p", 'app')

    checkError(output,error,"not_enough_args")

    assert_true status.success?, 'Process did not exit cleanly'
    assert_include output, "usage:", 'return was not correct'
  end

  def test_no_template
    output, error, status = Open3.capture3(PATH, BIN,"-s=\"showver.sh\"", "-t=\"no_template\"","-p", "app")

    checkNoError(output,error,"no_template")

    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, "exit status 2", 'wrong error'
  end

  def test_malformed_template
    output, error, status = Open3.capture3(PATH, BIN,"-s=\"showver.sh\"", "-t=\"useless_template\"","-p", "app")

    checkNoError(output,error,"malformed_template")

    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, "exit status 2", 'wrong error'
  end

  def test_copy_failed
    output, error, status = Open3.capture3(PATH, BIN, "-c=\"touch test && cp test ./test/test\"","-p", "app")

    checkNoError(output,error,"copy_failed")

    assert_false status.success?, 'Process did exit cleanly'
    assert_include error, "Process exited with status 1", 'wrong error'
  end

end



def checkError(output,error,testName)
  return if error.empty?
  puts "test: #{testName}"
  puts "output: #{output}"
  puts "error: #{error}"
end

def checkNoError(output,error,testName)
  return unless error.empty?
  puts "test: #{testName}"
  puts "output: #{output}"
  puts "error: #{error.inspect}"
end
