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

require 'rubygems'
require 'os'

desc 'generate a release tarball'
task release: [:compile] do
  release_dir = "#{release_path}/v#{APP_VERSION}"
  latest_dir  = "#{release_path}/latest"

  mkdir_p(release_dir)
  rm_rf(latest_dir)
  mkdir_p(latest_dir)

  cd(release_dir) { sh "tar czf #{APP_NAME}-#{APP_VERSION}.tgz #{src_path}" }

  Go::Build.builds.each do |gb|
    bin_path = "#{build_path}/#{gb.name}/bin/"
    tar_path = "#{APP_NAME}-#{APP_VERSION}-#{gb.name}.tgz"
    cd(release_dir) { sh "tar czf #{tar_path} #{bin_path}" }
  end

  ln Dir.glob("#{release_dir}/*"), latest_dir
end
