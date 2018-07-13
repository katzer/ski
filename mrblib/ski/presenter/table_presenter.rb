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
    # The default table columns
    COLUMNS = %w[NR. ID TYPE CONNECTION NAME OUTPUT].freeze
    # The default table style
    STYLE   = { all_separators: true }.freeze

    # Format and print the results to STDOUT.
    #
    # @param [ Array<SKI::Result> ] results The results to print out.
    #
    # @return [ Void ]
    def present(results)
      STDOUT.puts table(results).to_s
    end

    private

    # Internal table object.
    #
    # @param [ Array<SKI::Result> ] results The results to print out.
    #
    # @return [ Terminal::Table ]
    def table(results)
      Terminal::Table.new do |t|
        t.title    = ARGV.join(' ').sub(/^(.*?)(?=ski)/, '')
        t.headings = COLUMNS
        t.style    = STYLE
        t.rows     = rows(results)
        t.align_column 0, :right
      end
    end

    # Converts the results into table row structures.
    #
    # @param [ Array<SKI::Result> ] results The results to convert into rows.
    #
    # @return [ Array ]
    def rows(results)
      results.each_with_index { |res, idx| results[idx] = row(res, idx + 1) }
    end

    # Convert the results into table row structure.
    #
    # @param [ Result ] res The result to convert into a row.
    # @param [ Int ]    idx The index of the row to render.
    #
    # @return [ Array ]
    def row(res, idx)
      [
        "#{idx}.",
        res.planet.id,
        res.planet.type,
        res.planet.connection,
        res.planet.name,
        colorize_text(adjust(res), res.success)
      ]
    end
  end
end
