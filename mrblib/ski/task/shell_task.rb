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
  # Execute HTTP requests
  class WebTask < ServerTask
    # Request the REST-API of the web server.
    #
    # @param [ SKI::Planet ] planet The planet where to execute the task.
    #
    # @return [ Void ]
    def exec(planet)
      log("Executing shell command for #{planet.connection}") { shell(planet) }
    end

    # The command to execute on the remote server.
    #
    # @param [ SKI::Planet ] planet The planet where to execute the task.
    #
    # @return [ String ]
    def command(planet)
      [
        "export ORBIT_PLANET_ID=#{planet.id}",
        "export ORBIT_PLANET_URL=#{planet.connection}",
        super()
      ].join(";")
    end

    private

    # Execute the shell command for the remote server and yields the code block
    # with the captured result.
    #
    # @param [ SKI::Planet ] planet The planet where to execute the task.
    #
    # @return [ SKI::Result ]
    def shell(planet)
      out = `#{command(planet)}`

      result(planet, out, $? == 0)
    end
  end
end