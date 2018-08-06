# Apache 2.0 License
#
# Copyright (c) 2016 Sebastian Katzer, appPlant GmbH
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

require 'open3'
require_relative '../mrblib/ski/version'

BIN = File.expand_path('../mruby/bin/ski', __dir__).freeze

NO_KEY  = ENV.to_h.merge('ORBIT_KEY' => nil).freeze
BAD_KEY = ENV.to_h.merge('ORBIT_KEY' => 'bad file').freeze

assert('version [-v]') do
  output, status = Open3.capture2(BIN, '-v')

  assert_true status.success?, 'Process did not exit cleanly'
  assert_include output, SKI::VERSION
end

assert('version [--version]') do
  output, status = Open3.capture2(BIN, '--version')

  assert_true status.success?, 'Process did not exit cleanly'
  assert_include output, SKI::VERSION
end

assert('usage [-h]') do
  output, status = Open3.capture2(BIN, '-h')

  assert_true status.success?, 'Process did not exit cleanly'
  assert_include output, 'Usage'
end

assert('usage [--help]') do
  output, status = Open3.capture2(BIN, '--help')

  assert_true status.success?, 'Process did not exit cleanly'
  assert_include output, 'Usage'
end

assert('no $ORBIT_KEY') do
  _, output, status = Open3.capture3(NO_KEY, BIN, '-c', 'echo', 'host')

  assert_false status.success?, 'Process did exit cleanly'
  assert_include output, 'not set'
end

assert('bad $ORBIT_KEY') do
  _, output, status = Open3.capture3(BAD_KEY, BIN, '-c', 'echo', 'host')

  assert_false status.success?, 'Process did exit cleanly'
  assert_include output, 'not found'
end

assert('unknown flag') do
  _, output, status = Open3.capture3(BIN, '-unknown')

  assert_false status.success?, 'Process did exit cleanly'
  assert_include output, 'unknown option'
end

assert('no command') do
  _, output, status = Open3.capture3(BIN, 'host')

  assert_false status.success?, 'Process did exit cleanly'
  assert_include output, 'ArgumentError'
end

assert('command and script') do
  _, output, status = Open3.capture3(BIN, '-c', 'echo', '-s', 'path', 'host')

  assert_false status.success?, 'Process did exit cleanly'
  assert_include output, 'ArgumentError'
end

assert('pretty [-p]') do
  skip if ENV['OS'] == 'Windows_NT'

  output, status = Open3.capture2(BIN, '-p', '-c', 'echo test', 'server')

  assert_true status.success?, 'Process did not exit cleanly'
  assert_include output, '| NR. | ID          | NAME   | OUTPUT                      |'
  assert_include output, "| test                        |\n"
end

assert('no color [--no-color]') do
  skip if ENV['OS'] == 'Windows_NT'

  output, status = Open3.capture2(BIN, '--no-color', '--pretty', '-c', 'echo test', 'server')

  assert_true status.success?, 'Process did not exit cleanly'
  assert_include output, "| ArgumentError: 'initialize' |\n"
end

assert('width [-w]') do
  skip if ENV['OS'] == 'Windows_NT'

  output, status = Open3.capture2(BIN, '-p', '-w', '6', '-c', '123', 'server')

  assert_true status.success?, 'Process did not exit cleanly'
  assert_include output, '| NR. | ID          | NAME   | OUTPUT |'
end

assert('no matcher') do
  _, output, status = Open3.capture3(BIN, '-c', 'echo')

  assert_false status.success?, 'Process did exit cleanly'
  assert_include output, 'ArgumentError'
end

assert('bad script') do
  _, output, status = Open3.capture3(BIN, '-s', 'path', 'host')

  assert_false status.success?, 'Process did exit cleanly'
  assert_include output, 'NoFileError'
end

assert('bad template') do
  _, output, status = Open3.capture3(BIN, '-c', 'echo', '-t', 'path', 'host')

  assert_false status.success?, 'Process did exit cleanly'
  assert_include output, 'NoFileError'
end
