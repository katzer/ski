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
  # Print the output to STDOUT as a table
  class TablePresenter < BasePresenter
    # initialize the table presenter object.
    #
    # @return [ Void ]
    def initialize(*)
      @title   = ARGV.join(' ').sub(/^(.*?)(?=ski)/, '')
      @columns = %w[NR. ID TYPE CONNECTION NAME OUTPUT]
      @style   = { all_separators: true }
      super
    end

    # Format and print the results to STDOUT.
    #
    # @param [ Array<SKI::Result> ] *results 1 to n results to print out.
    #
    # @return [ Void ]
    def print(*results)
      STDOUT.puts table(*results).to_s
    end

    private

    # Internal table object.
    #
    # @param [ Array<SKI::Result> ] *results 1 to n results to print out.
    #
    # @return [ Terminal::Table ]
    def table(*results)
      Terminal::Table.new do |t|
        t.headings = @columns
        t.style    = @style
        t.rows     = rows(*results)
        t.title    = @title
        t.align_column 0, :right
      end
    end

    # Converts the results into table row structures.
    #
    # @param [ Array<SKI::Result> ] results The results to convert into rows.
    #
    # @return [ Array ]
    def rows(*results)
      pos = 0

      results.map! do |r, p = r.planet|
        ["#{pos += 1}.", p.id, p.type, p.connection, p.name, colorize_output(r)]
      end
    end
  end
end
