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
  # Base class for all task types
  class BaseTask
    # Default configuration for every SSH connection
    SSH_CONFIG = { key: ENV['ORBIT_KEY'], compress: true, timeout: 5000 }.freeze

    # Initialize the task specified by opts.
    #
    # @param [ Hash<Symbol,Object> ] opts A key-value hash.
    #
    # @return [ Void ]
    def initialize(opts)
      @opts = opts
    end

    protected

    # The command to execute on the remote server.
    #
    # @return [ String ]
    def command
      (@opts[:script] ? IO.read(@opts[:script]) : @opts[:command])&.strip
    end

    # Return a task result.
    #
    # @param [ SKI::Planet ] planet   The planet where the result comes from.
    # @param [ String ]      output   The output of the task execution.
    # @param [ Boolean ]     no_error If the task executed without an error.
    #                                 Defaults to: true
    #
    # @return [ SKI::Result ]
    def result(planet, output, no_error = true)
      output = %("#{output.chomp! || output}") if @opts[:template] && !no_error
      Result.new(planet, output, no_error)
    end

    # Return a task result where the error bit is set.
    #
    # @param [ SKI::Planet ] planet The planet where the result comes from.
    # @param [ String ]      output The output of the task execution.
    #
    # @return [ SKI::Result ]
    def error(planet, output)
      logger.error(output)
      result(planet, output, false)
    end

    private

    # Logging device that writes into $ORBIT_HOME/log/plip.log
    #
    # @return [ Logger ]
    def logger
      $logger ||= begin
        dir = File.join(ENV['ORBIT_HOME'], 'logs')
        Dir.mkdir(dir) unless Dir.exist? dir

        Logger.new("#{dir}/ski.log", formatter: lambda do |sev, ts, _, msg|
          "[#{sev[0, 3]}] #{ts}: #{msg}\n"
        end)
      end
    end

    # Write a log message, execute the code block and write another log.
    # that the task is done.
    #
    # @param [ String ] msg The message to log.
    # @param [ Proc ] block The code block to execute.
    #
    # @return [ Void ]
    def log(msg)
      logger.info msg
      res = yield
      logger.info "#{msg} done"
      res
    end

    # Write an error log message.
    #
    # @param [ String ] user The remote user.
    # @param [ String ] host The remote host.
    # @param [ SSH::Session ] ssh The connected SSH session.
    # @param [ String ] msg  The error message.
    #
    # @return [ Void ]
    def log_error(usr, host, ssh, msg = nil)
      logger.error "#{usr}@#{host} #{ssh&.last_error} #{ssh&.last_errno} #{msg}"
    end

    # Establish a SSH connection to the host and make sure that the timeout is
    # only used for the connect period.
    #
    # @param [ String ] user The remote user.
    # @param [ String ] host The remote host.
    #
    # @return [ SSH::Session ]
    def __connect__(user, host)
      log "Connecting to #{user}@#{host}" do
        ssh = SSH.start(host, user, SSH_CONFIG.dup)
        ssh.timeout = 0
        ssh
      end
    end

    # Start an SSH session.
    #
    # @param [ SKI::Planet ] planet The planet where to connect to.
    #
    # @return [ Void ]
    def connect(planet)
      user, host = planet.user_and_host
      ssh        = __connect__(user, host)
      res        = yield(ssh, planet)
      log_error(user, host, ssh) if ssh.last_error
      res
    rescue RuntimeError => e
      log_error(user, host, ssh, e.message)
      result(planet, e.message, false)
    ensure
      ssh&.close
    end
  end
end
