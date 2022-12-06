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

# codebeat:disable[TOO_MANY_IVARS]

module SKI
  # Wrapper around the planet details that fifa returned in ski format
  class Planet < BasicObject
    # Initialize a wrapper for a planet instance.
    #
    # @param [ String ] res The string returned by "fifa -f ski ..."
    #
    # @return [ Void ]
    def initialize(res)
      @suc, @id, @type, @name, @connection = res.split('|')
      @suc                                 = @suc == '1'
    end

    attr_reader :id, :type, :name, :connection

    # If the planet contains valid information.
    #
    # @return [ Boolean ]
    def valid?
      @suc
    end

    # The id of the database.
    #
    # @return [ String ]
    def db
      @connection.split(':')[0] if valid? && db?
    end

    # The name of the SSH user.
    #
    # @return [ String ]
    def user
      @connection.split('@')[0].split(':')[-1] if valid?
    end

    # The user and host of the SSH server.
    #
    # @return [ Array<String> ]
    def user_and_host
      user, host = @connection.split('@') if valid?
      user       = user.split(':')[-1]    if user && db?

      [user, host]
    end

    # Execute the task depend on the type of the planet.
    #
    # @param [ Hash ] spec The spec for the task.
    #
    # @return [ SKI::Result ] The result of the task.
    def exec(spec)
      case task_type(spec)
      when 'server'  then ServerTask.new(spec)
      when 'db'      then DatabaseTask.new(spec)
      when :error    then ErrorTask.new(spec)
      when :invalid  then InvalidTask.new(spec)
      else                ShellTask.new(spec)
      end.exec(self)
    end

    private

    # Test if the planet is from type 'db'.
    #
    # @return [ Boolean ]
    def db?
      type == 'db'
    end

    # Internal helper method that determines the right task class based on what
    # fifa responded and if the planet type and script extension match.
    #
    # @param [ Hash ] spec The spec for the task.
    #
    # @return [ Boolean ]
    def task_type(spec)
      return :error unless valid?

      file = spec[:script]

      return :invalid if file && ((db? && file !~ /\.sql$/) || \
                                 (!db? && file !~ /\.sh$/))

      type
    end
  end
end

# codebeat:enable[TOO_MANY_IVARS]
