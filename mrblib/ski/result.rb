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
  # Encapsulate the result of a task execution.
  class Result
    # Create a result object with contains all infos about a task result.
    #
    # @param [ SKI::Planet ] planet The planet where the task has been executed.
    # @param [ String ]      output The output of the remote execution.
    # @param [ Boolean ]    success If the task executed in a successful way.
    #
    # @return [ Void ]
    def initialize(planet, output, success)
      @planet     = planet
      @success    = success
      self.output = output
    end

    # The planet where the task has been executed.
    #
    # @return [ SKI::Planet ]
    attr_reader :planet

    # The output of the remote execution.
    #
    # @return [ String ]
    attr_reader :output

    # If the task executed in a successful way.
    #
    # @return [ Boolean ]
    attr_accessor :success

    # Setter for the result output.
    #
    # @param [ String ] output The output of the remote execution.
    #
    # @return [ String ]
    def output=(output)
      @output = output&.chop! || output || ''
    end
  end
end
