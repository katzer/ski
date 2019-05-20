# Apache 2.0 License
#
# Copyright (c) 2018 Sebastian Katzer, appPlant GmbH
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

ENV['MRUBY_CLI_LOCAL'] ||= '1'   if ENV['OS'] == 'Windows_NT'
ENV['TOOLCHAIN']       ||= 'gcc' if ENV['OS'] == 'Windows_NT'

def in_a_docker_container?
  Open3.capture2e('grep -q docker /proc/self/cgroup')[-1].success?
end

def docker_run(cmd, ver = ENV['MRUBY_VERSION'])
  sh "MRUBY_VERSION=#{ver} docker-compose run #{cmd}"
end

Dir["#{__dir__}/tasks/**/*.rake"].each { |file| load file }
