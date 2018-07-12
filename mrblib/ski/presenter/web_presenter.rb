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
  # Save the output under /results/<job> for iss
  class WebPresenter < BasicObject
    # Absolute path to $ORBIT_HOME/reports
    REPORTS_DIR = File.join(ENV['ORBIT_HOME'].to_s, 'reports').freeze

    # Initialize a new presenter object.
    #
    # @param [ Boolean ] job  The name of the job.
    # @param [ String ]  cols The name of the columns.
    #
    # @return [ Void ]
    def initialize(job, cols)
      @job  = job
      @ts   = Time.now.to_i
      @cols = columns(cols || 'OUTPUT')
    end

    # Format and print the results to $ORBIT_HOME/reports.
    #
    # @param [ Array<SKI::Result> ] results The results to print out.
    #
    # @return [ Void ]
    def present(results)
      open_report_file do |f|
        f << "#{@ts}\n#{@cols}"

        results.each do |r, p = r.planet|
          r.output.split("\n").each do |o|
            f << %(\n["#{p.id}","#{p.name}",#{r.success},#{o}])
          end
        end
      end
    end

    private

    # Convert the columns into tuples of name and type.
    #
    # @param [ String ] The columns to convert.
    #
    # @return [ String ]
    def columns(cols)
      cols.split.map! do |name|
        case name[-2, 2]
        when '_s' then [name[0...-2], 'string']
        when '_i' then [name[0...-2], 'int']
        when '_f' then [name[0...-2], 'float']
        else           [name,         'string']
        end
      end.inspect
    end

    # Open the file and pass to the code block to invoke.
    #
    # @param [ Proc ] &block The code block to call for.
    #
    # @return [ Void ]
    def open_report_file
      File.open("#{make_report_file_dir}/#{@ts}.skirep", 'w+') do |f|
        yield(f) && STDOUT.puts("Written report to: #{f.path}")
      end
    end

    # Create all parent directories within $ORBIT_HOME
    #
    # @return [ String ] The report sub folder
    def make_report_file_dir
      rep_dir = File.join(REPORTS_DIR, @job)

      [REPORTS_DIR, rep_dir].each { |dir| Dir.mkdir(dir) unless Dir.exist? dir }

      rep_dir
    end
  end
end
