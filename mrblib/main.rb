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

@parser = OptParser.new do |opts|
  opts.add :command,    :string
  opts.add :script,     :string
  opts.add :job,        :string
  opts.add :template,   :string
  opts.add :width,      :int, 0
  opts.add :pretty,     :bool, false
  opts.add :'no-color', :bool, false
end

@parser.on! :help do
  <<-USAGE
Usage: ski [options...] matchers...
Options:
-c, --command   Execute command and return result
-s, --script    Execute script and return result
-t, --template  Template to be used to transform the output
-j, --job       Execute job specified in file
-n, --no-color  Print errors without colors
-p, --pretty    Pretty print output as a table
-w, --width     Width of output column in characters
-h, --help      This help text
-v, --version   Show version number
USAGE
end

@parser.on! :version do
  "ski #{SKI::VERSION} - #{OS.sysname} #{OS.bits(:binary)}-Bit (#{OS.machine})"
end

# Entry point of the tool.
#
# @param [ Array<String> ] args The ARGV array.
#
# @return [ Void ]
def __main__(args)
  SKI::Job.new(__parse__(args[1..-1])).exec
end

# Parse the command-line arguments.
#
# @param [ Array<String> ] args The command-line arguments.
#
# @return [ Hash<Symbol,Object> ]
def __parse__(args)
  opts        = @parser.parse(args.empty? ? ['-h'] : args)
  opts        = __parse_skijob__(opts[:job]) if opts[:job]
  opts[:tail] = @parser.tail
  opts
end

# Parse the skijob file.
#
# @param [ String ] job The name of the file.
#
# @return [ Hash<Symbol,Object> ]
def __parse_skijob__(job)
  @parser.parse IO.read("#{ENV['ORBIT_HOME']}/job/#{job}.skijob")
                  .split(/(?<!\\)\s+/)
                  .map! { |f| f.gsub(/[\\'"]/, '') }
                  .concat(['-j', job])
end
