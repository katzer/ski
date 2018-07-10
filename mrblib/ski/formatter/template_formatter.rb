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
  # Convert the output based on a TextFSM template file.
  class TemplateFormatter
    # The absolute path to the skifsm.pyc script
    FSM_PATH = File.join(ENV['ORBIT_HOME'].to_s, 'vendor/textfsm/skifsm.py')

    # Initialize a new formatter.
    #
    # @param [ String ] tpl The path of the template file.
    #
    # @return [ Void ]
    def initialize(tpl)
      @tpl = tpl
    end

    # Replace output for each result through parsed textfsm template.
    #
    # @param [ Array<SKI::Result> ] results The results to format.
    #
    # @return [ Array<SKI::Result> ]
    def format(results)
      results.each { |res| update_output_by_template(res) if res.success }
      results
    end

    # The columns for the template.
    #
    # @return [ String ]
    def columns
      skifsm('--columns').chomp!
    end

    # Update output for each result through parsed textfsm template.
    #
    # @param [ SKI::Result ] res The result to update.
    #
    # @return [ SKI::Result ]
    def update_output_by_template(res)
      res.output  = skifsm("<<EOF\n#{res.output}\nEOF")
      res.success = $? == 0
    end

    # Invoke the skifsm.pyc script with the given additional arguments.
    #
    # @param [ String ] args The additional arguments.
    #
    # return [ String ]
    def skifsm(args)
      `python #{FSM_PATH} #{@tpl} #{args}`
    end
  end
end
