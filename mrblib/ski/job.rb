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
  # Coordinates the parallel execution of tasks and result aggregation.
  class Job
    # Initialize a new job that coordinates the execution of the tasks.
    #
    # @param [ Hash<Symbol,Object> ] spec The parsed command line arguments.
    #
    # @return [ Void ]
    def initialize(spec)
      @spec = convert(spec)
    end

    # Execute the command specified by the opts.
    #
    # @return [ Void ]
    def exec
      if @spec[:pretty] || @spec[:job]
        exec_and_present_aggregated
      else
        exec_and_present
      end
    end

    private

    # Split the job into multiple tasks, execute them
    # and present the result once a job is done.
    #
    # @return [ Void ]
    def exec_and_present
      validate && async { |opts| present([SKI::Planet.new(opts).exec(@spec)]) }
    end

    # Split the job into multiple tasks, execute them
    # and present all results at once when all jobs are done.
    #
    # @return [ Void ]
    def exec_and_present_aggregated
      validate && present(async { |opts| SKI::Planet.new(opts).exec(@spec) })
    end

    # Devide the list of planets into slices and execute
    # the code block for each slice within an own thread.
    #
    # @param [ Proc ] &block A code block to execute per slice.
    #
    # @return [ Array<SKI::Result> ]
    def async(&block)
      servers = planets
      size    = [(servers.count / 20.0).round, 1].max
      ths     = []

      servers.each_slice(size) do |slice|
        ths << Thread.new(slice) { |opts| opts.map! { |opt| block&.call(opt) } }
      end

      ths.map!(&:join).flatten!&.compact!
      ths
    end

    # Convert the parsed spec values.
    #
    # @param [ Hash<Symbol,Object> ] opt A key:value map.
    #
    # @return [ Hash<Symbol,Object> ] opt
    def convert(opt)
      script, tpl = opt[:script], opt[:template]

      opt[:template] = expand_path('template', "#{tpl}.textfsm") if tpl
      opt[:script]   = expand_path('script', script) if script&.include? '.sh'
      opt[:script]   = expand_path('sql', script)    if script&.include? '.sql'

      opt
    end

    # Validate the parsed command-line arguments.
    # Raises an error in case of something is missing or invalid.
    #
    # @return [ Boolean ] true if valid
    def validate
      validate_envs && validate_args
    end

    # rubocop:disable AbcSize, CyclomaticComplexity, LineLength, PerceivedComplexity

    # Validate the parsed command-line arguments.
    # Raises an error in case of something is missing or invalid.
    #
    # @return [ Boolean ] true if valid
    def validate_args
      raise ArgumentError,     'Missing command or script'          unless @spec[:command] || @spec[:script]
      raise ArgumentError,     'Execute with command or script'     if     @spec[:command] && @spec[:script]
      raise ArgumentError,     'Missing matcher'                    unless @spec[:tail].any?
      raise File::NoFileError, "No such file - #{@spec[:script]}"   if     @spec[:script]   && !File.file?(@spec[:script])
      raise File::NoFileError, "No such file - #{@spec[:template]}" if     @spec[:template] && !File.file?(@spec[:template])

      true
    end

    # Validate environment arguments.
    # Raises an error in case of something is missing or invalid.
    #
    # @return [ Boolean ] true if valid
    def validate_envs
      raise KeyError,          '$ORBIT_KEY not set'   unless ENV['ORBIT_KEY']
      raise File::NoFileError, '$ORBIT_KEY not found' unless File.exist? ENV['ORBIT_KEY']

      true
    end

    # rubocop:enable AbcSize, CyclomaticComplexity, LineLength, PerceivedComplexity

    # Retrieve list of servers in ski format from fifa.
    #
    # @return [ Array<String> ]
    def planets
      query = @spec[:tail].join('" "')
      fifa  = ENV.include?('ORBIT_BIN') ? "#{ENV['ORBIT_BIN']}/fifa" : 'fifa'
      cmd   = %(#{fifa} -n -f ski "#{query}")
      out   = `#{cmd}`

      raise "#{cmd} failed with exit code #{$?}" unless $? == 0

      out.split("\n").map!(&:chomp)
    end

    # Helper to construct absolute path.
    #
    # @param [ Array<String> ] *folders The folders to join with $ORBIT_HOME.
    #
    # @return [ String ]
    def expand_path(*folders)
      File.join(ENV.fetch('ORBIT_HOME'), *folders)
    rescue KeyError
      raise '$ORBIT_HOME not set'
    end

    # The formatter for the output.
    #
    # @return [ SKI::BaseFormatter ]
    def formatter
      @formatter ||= TemplateFormatter.new(@spec[:template]) if @spec[:template]
    end

    # The presenter for the output.
    #
    # @return [ SKI::BasePresenter ]
    def presenter
      @presenter ||= if @spec[:job]
                       WebPresenter.new(@spec, formatter&.columns)
                     elsif @spec[:pretty]
                       TablePresenter.new(@spec, formatter&.columns)
                     else
                       PlainPresenter.new(@spec)
                     end
    end

    # Format the results and send them to the presenter.
    #
    # @param [ Array<SKI::Result> ] *results The results to format and present.
    #
    # @return [ Void ]
    def present(results)
      results = formatter.format(results) if formatter
      presenter.present(results) && nil
    end
  end
end
