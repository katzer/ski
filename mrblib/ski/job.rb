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
    # @param [ Hash ] spec The parsed command line arguments.
    #
    # @return [ Void ]
    def initialize(spec)
      @spec = validate convert spec
    end

    # Split the job into multiple tasks and execute them asynchron.
    #
    # @param [ Hash<Symbol, _> ] opts A key:value map.
    #
    # @return [ Void ]
    def exec
      case presenter
      when TablePresenter
        present(*async { |opts| SKI::Planet.new(*opts).exec(@spec) })
      when PlainPresenter
        async { |opts| present SKI::Planet.new(*opts).exec(@spec) }
      end
    end

    private

    # Devide the list of planets into slices and execute the code block
    # for each slice within an own thread.
    #
    # @param [ Proc ] &block A code block to execute per slice.
    #
    # @return [ Void ]
    def async(&block)
      servers = planets
      size    = [(servers.count / 20.0).round, 1].max
      ths     = []

      servers.each_slice(size) do |slice|
        ths << Thread.new(slice) { |opts| opts.map! { |opt| block&.call(opt) } }
      end

      ths.map!(&:join).each_with_object([]) { |res, o| o.concat(res) if res[0] }
    end

    # Convert the parsed spec values.
    #
    # @param [ Hash<Symbol,Object> ] opt A key:value map.
    #
    # @return [ Hash<Symbol,Object> ] opt
    def convert(opt)
      script = opt[:script]

      opt[:template] = expand_path('template', opt[:template]) if opt[:template]
      opt[:script]   = expand_path('scripts', script) if script&.include? '.sh'
      opt[:script]   = expand_path('sql', script)     if script&.include? '.sql'

      opt
    end

    # rubocop:disable AbcSize, CyclomaticComplexity, LineLength, PerceivedComplexity

    # Validate the parsed command-line arguments.
    # Raises an error in case of something is missing or invalid.
    #
    # @param [ Hash<Symbol,Object> ] opt A key:value map.
    #
    # @return [ Hash<Symbol,Object> ] opt
    def validate(opt)
      raise ArgumentError,     'Missing command or script'        unless opt[:command] || opt[:script]
      raise ArgumentError,     'Execute with command or script'   if     opt[:command] && opt[:script]
      raise ArgumentError,     'Missing matcher'                  unless opt[:tail].any?
      raise File::NoFileError, "No such file - #{opt[:script]}"   if     opt[:script]   && !File.file?(opt[:script])
      raise File::NoFileError, "No such file - #{opt[:template]}" if     opt[:template] && !File.file?(opt[:template])

      opt
    end

    # rubocop:enable AbcSize, CyclomaticComplexity, LineLength, PerceivedComplexity

    # Server list retrieved from fifa.
    #
    # @return [ Array<"user@host"> ]
    def planets
      cmd = %(#{ENV['ORBIT_BIN']}/fifa -f=ski "#{@spec[:tail].join('" "')}")
      out = `#{cmd}`

      raise "#{cmd} failed with exit code #{$?}" unless $? == 0

      out.split("\n").map! { |line| line.split('|') }
    end

    # Helper to construct absolute path.
    #
    # @param [ Array<String> ] *folders The folders to join with $ORBIT_HOME.
    #
    # @return [ String ]
    def expand_path(*folders)
      File.join(ENV['ORBIT_HOME'], *folders)
    end

    # The presenter for the job output.
    #
    # @return [ SKI::BasePresenter ]
    def presenter
      @presenter ||= begin
        if @spec[:pretty]
          TablePresenter.new(@spec[:'no-color'])
        else
          PlainPresenter.new(@spec[:'no-color'])
        end
      end
    end

    # Format the results and send them to the presenter.
    #
    # @param [ Array<SKI::Result> ] *results The results to format and present.
    #
    # @return [ Void ]
    def present(*results)
      presenter.print(*results) && nil
    end
  end
end
