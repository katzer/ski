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

require 'os'
require 'fileutils'

require_relative 'src/tasks/build.rb'
require_relative 'build_config.rb'

APP_NAME     = ENV['APP_NAME'] || 'goo'
APP_ROOT     = ENV['APP_ROOT'] || Dir.pwd
APP_VERSION  = ENV['APP_VERSION'] || '0.0.1'

release_path = "#{APP_ROOT}/releases"
build_path   = "#{APP_ROOT}/build"
src_path     = "#{APP_ROOT}/src"

desc 'compile binary'
task :compile do
  Goo::Build.builds.each do |gb|
    bin_path = "#{build_path}/#{gb.name}/bin"
    goo_path = "#{src_path}/#{APP_NAME}.go"
    mkdir_p(bin_path)

    cd(bin_path) do
      if OS.windows?
        sh "set GOOS=#{gb.os}&&set GOARCH=#{gb.arch}&&go build #{goo_path}"
      else
        sh "GOOS=#{gb.os} GOARCH=#{gb.arch} go build #{goo_path}"
      end
    end
  end
end

namespace :test do
  desc 'run integration tests'
  task :bintest do
    Goo::Build.builds.each do |gb|
      next unless gb.bintest?
      bin_path = "#{build_path}/#{gb.name}/bin/goo"
      bin_path << '.exe' if File.exist? "#{bin_path}.exe"
      sh "ruby #{APP_ROOT}/bintest/goo.rb #{bin_path}"
    end
  end
end

desc 'cleanup builds'
task :clean do
  rm_rf build_path
end

desc 'generate a release tarball'
task release: :compile do
  release_dir = "#{release_path}/v#{APP_VERSION}"
  latest_dir  = "#{release_path}/latest"

  mkdir_p(release_dir)
  rm_rf(latest_dir)
  mkdir_p(latest_dir)
  cd(release_dir) { sh "tar czf #{APP_NAME}-#{APP_VERSION}.tgz #{src_path}" }

  Goo::Build.builds.each do |gb|
    bin_path = "#{build_path}/#{gb.name}/bin/"
    tar_path = "#{APP_NAME}-#{APP_VERSION}-#{gb.name}.tgz"
    cd(release_dir) { sh "tar czf #{tar_path} #{bin_path}" }
  end

  ln Dir.glob("#{release_dir}/*"), latest_dir
end

task :version do
  puts `#{build_path}/Linux64/goo -v`
end
