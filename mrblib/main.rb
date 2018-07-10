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
  opts.add :pretty,     :bool, false
  opts.add :'no-color', :bool, false
end

@parser.on! :help do
  <<-USAGE

#{SKI::LOGO}

usage: ski [options...] matchers...
Options:
-c, --command   Execute command and return result
-s, --script    Execute script and return result
-t, --template  Template to be used to transform the output
-j, --job       Execute job specified in file
-p, --pretty    Pretty print output as a table
--no-color      Print errors without colors
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
  validate && SKI::Job.new(parse(args[1..-1])).exec
end

# Parse the command-line arguments.
#
# @param [ Array<String> ] args The command-line arguments to parse.
#
# @return [ Hash<Symbol, Object> ]
def parse(args)
  opts = @parser.parse(args.empty? ? ['-h'] : args)

  if opts[:job]
    args = File.read("#{ENV['ORBIT_HOME']}/jobs/#{opts[:job]}.skijob")
               .split(/(?<!\\)\s+/).map! { |f| f.gsub(/[\\'"]/, '') }
    opts = @parser.parse(args)
  end

  opts[:tail] = @parser.tail
  opts
end

# Validate the environment variables.
# Raises an error in case of something is missing or invalid.
#
# @return [ Void ]
def validate
  %w[ORBIT_HOME ORBIT_BIN ORBIT_KEY].each do |env|
    raise "#{env} not set"   unless ENV[env]
    raise "#{env} not found" unless File.exist? ENV[env]
  end
end
