#
# Copyright (c) 2013-2016 by appPlant GmbH. All rights reserved.
#
# @APPPLANT_LICENSE_HEADER_START@
#
# This file contains Original Code and/or Modifications of Original Code
# as defined in and that are subject to the Apache License
# Version 2.0 (the 'License'). You may not use this file except in
# compliance with the License. Please obtain a copy of the License at
# http://opensource.org/licenses/Apache-2.0/ and read it before using this
# file.
#
# The Original Code and all software distributed under the License are
# distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
# EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
# INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
# Please see the License for the specific language governing rights and
# limitations under the License.
#
# @APPPLANT_LICENSE_HEADER_END@

module Go
  # Cross build specification for Golang
  class Build
    attr_reader :name

    def self.builds
      @builds ||= []
    end

    def initialize(name, &block)
      @name  = name
      @attrs = {}

      instance_exec(&block)
      self.class.builds << self
    end

    def os(name = nil)
      @attrs[:os] = name if name
      @attrs[:os]
    end

    def arch(name = nil)
      @attrs[:arch] = name if name
      @attrs[:arch]
    end

    def appname(name = nil)
      @attrs[:appname] = name if name
      @attrs[:appname]
    end

    def bintest_if(enabled)
      @attrs[:test] = enabled
      @attrs[:test]
    end

    def bintest?
      @attrs[:test] == true
    end

    def go_build(binpath)
      %(#{env_setup} go build -i -ldflags="-s -X #{version_cmd}" -o #{binpath}/#{appname})
    end

    private

    def env_setup
      if OS.windows?
        "set GOOS=#{os} && set GOARCH=#{arch} &&"
      else
        "GOOS=#{os} GOARCH=#{arch}"
      end
    end

    def version_cmd
      "main.version=$(go run #{version_path}/*.go)"
    end
  end
end
