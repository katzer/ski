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

module SKI
  class Job
    # Initialize a new job that coordinates the execution of the tasks.
    #
    # @param [ Hash ] spec The parsed command line arguments.
    #
    # @return [ Void ]
    def initialize(spec)
      @spec = convert spec
    end

    # Download/Upload the file specified by the opts.
    #
    # @param [ Hash<Symbol, _> ] opts A key:value map.
    #
    # @return [ Void ]
    def exec
      validate
    end

    private

    # rubocop:disable AbcSize, CyclomaticComplexity, LineLength, PerceivedComplexity

    # Convert the parsed spec values.
    #
    # @param [ Hash<Symbol,Object> ] opts A key:value map.
    #
    # @return [ Hash<Symbol,Object> ] opts
    def convert(opts)
      opts[:job]      = "#{ENV['ORBIT_HOME']}/jobs/#{opts[:job]}.json"      if opts[:job]
      opts[:template] = "#{ENV['ORBIT_HOME']}/templates/#{opts[:template]}" if opts[:template]

      if opts[:script]&.include?('.sh')
        opts[:script] = "#{ENV['ORBIT_HOME']}/scripts/#{opts[:script]}"
      elsif opts[:script]&.include?('.sql')
        opts[:script] = "#{ENV['ORBIT_HOME']}/sql/#{opts[:script]}"
      end

      opts
    end

    # Validate the parsed command-line arguments.
    # Raises an error in case of something is missing or invalid.
    #
    # @return [ Boolean ] true if valid
    def validate
      raise ArgumentError,     'Missing command or script'          unless @spec[:command] || @spec[:script]
      raise ArgumentError,     'Missing matcher'                    unless @spec[:tail].any?
      raise File::NoFileError, "No such file - #{@spec[:script]}"   if @spec[:script]   && !File.file?(@spec[:script])
      raise File::NoFileError, "No such file - #{@spec[:template]}" if @spec[:template] && !File.file?(@spec[:template])
      true
    end

    # rubocop:enable AbcSize, CyclomaticComplexity, LineLength, PerceivedComplexity

    # Server list retrieved from fifa.
    #
    # @return [ Array<"user@host"> ]
    def planets
      cmd = %(#{ENV['ORBIT_BIN']}/fifa -f=ssh "#{@spec[:tail].join('" "')}")
      out = `#{cmd}`

      raise "#{cmd} failed with exit code #{$?}" unless $? == 0

      out.split("\n").map! { |ssh| ssh.split('@') }
    end
  end
end
