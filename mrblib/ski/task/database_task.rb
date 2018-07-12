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
  # Execute SQL command on the remote database.
  class DatabaseTask < BaseTask
    # The shell command to invoke pqdb_sql.out
    PQDB = '. profiles/%s.prof > /dev/null && exe/pqdb_sql.out -s -x %s'.freeze

    # Execute the SQL command on the remote database.
    #
    # @param [ SKI::Planet ] planet The planet where to execute the task.
    #
    # @return [ Void ]
    def exec(planet)
      connect(planet) do |ssh|
        log "Executing SQL command on #{ssh.host}" do
          cmd = format(PQDB, planet.user, planet.db)
          pqdb(ssh, cmd) { |out, ok| result(planet, out, ok) }
        end
      end
    end

    protected

    # Well formatted SQL command to execute on the remote server.
    #
    # @return [ String ]
    def command
      (cmd = super)[-1] == ';' ? cmd : "#{cmd};"
    end

    private

    # Execute the SQL command on the remote database and yields the code
    # block with the captured result.
    #
    # @param [ SSH::Session ] ssh      The SSH session that is connected to
    #                                  the remote host.
    # @param [ String ]       pqdb_cmd The shell command to invoke pqdb_sql.
    #
    # @return [ SKI::Result ]
    def pqdb(ssh, pqdb_cmd, &block)
      io, ok = ssh.open_channel.popen2e(pqdb_cmd)

      io.puts(command)
      io.puts('exit')

      block&.call(io.gets(nil), ok)
    ensure
      io&.close(false)
    end
  end
end
