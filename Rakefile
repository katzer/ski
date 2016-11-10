require 'fileutils'
require 'os'

MRUBY_VERSION="1.2.0"

file :mruby do
  #sh "git clone --depth=1 https://github.com/mruby/mruby"
  sh "curl -L -L --fail --retry 3 --retry-delay 1 https://github.com/mruby/mruby/archive/1.2.0.tar.gz -s -o - | tar zxf -"
  FileUtils.mv("mruby-1.2.0", "mruby")
end

APP_NAME=ENV["APP_NAME"] || "goo"
APP_ROOT=ENV["APP_ROOT"] || Dir.pwd
bin_path="#{APP_ROOT}/bin"
tools_path="#{ENV["TOOLS_PATH"]}"

desc "compile binary"
task :compile do
  sh "rm -r tools" unless !Dir.exists?(tools_path)
  sh "mkdir tools" unless Dir.exists?(tools_path)
  sh "rm -r /go/bin" unless !Dir.exists?("/go/bin")
  sh "mkdir /go/bin"
  Dir.chdir("/go/bin")
  sh "mkdir Linux64"
  sh "mkdir Linuxi686"
  sh "mkdir Win64"
  sh "mkdir Wini686"
  sh "mkdir Mac64"
  sh "mkdir Maci386"
  Dir.chdir("/go/bin/Linux64")
  sh"GOOS=linux GOARCH=amd64 go build /go/src/goo.go"
  Dir.chdir("/go/bin/Linuxi686")
  sh"GOOS=linux GOARCH=386 go build /go/src/goo.go"
  Dir.chdir("/go/bin/Win64")
  sh"GOOS=windows GOARCH=amd64 go build /go/src/goo.go"
  Dir.chdir("/go/bin/Wini686")
  sh"GOOS=windows GOARCH=386 go build /go/src/goo.go"
  Dir.chdir("/go/bin/Mac64")
  sh"GOOS=darwin GOARCH=amd64 go build /go/src/goo.go"
  Dir.chdir("/go/bin/Maci386")
  sh"GOOS=darwin GOARCH=386 go build /go/src/goo.go"

  if OS.linux?
    if OS.bits == 64
      sh"GOOS=linux GOARCH=amd64 go build /go/src/goo.go"
      Dir.chdir(tools_path)
      sh "curl -L https://github.com/appPlant/ff/releases/download/#{ENV["FF_VER"]}/ff-#{ENV["FF_VER"]}-x86_64-pc-linux-gnu.tgz  | tar xz"
    elsif OS.bits == 32
      sh"GOOS=linux GOARCH=386 go build /go/src/goo.go"
      Dir.chdir(tools_path)
      sh "curl -L https://github.com/appPlant/ff/releases/download/#{ENV["FF_VER"]}/ff-#{ENV["FF_VER"]}-i686-pc-linux-gnu.tgz  | tar xz"
    end
  elsif OS.mac?
    if OS.bits == 64
      sh"GOOS=darwin GOARCH=amd64 go build /go/src/goo.go"
      Dir.chdir(tools_path)
      sh "curl -L https://github.com/appPlant/ff/releases/download/#{ENV["FF_VER"]}/ff-#{ENV["FF_VER"]}-x86_64-apple-darwin14.tgz  | tar xz"
    elsif OS.bits == 32
      sh"GOOS=darwin GOARCH=386 go build /go/src/goo.go"
      Dir.chdir(tools_path)
      sh "curl -L https://github.com/appPlant/ff/releases/download/#{ENV["FF_VER"]}/ff-#{ENV["FF_VER"]}-i386-apple-darwin14.tgz | tar xz"
    end
  elsif OS.windows?
    if OS.bits == 64
      sh"GOOS=windows GOARCH=amd64 go build /go/src/goo.go"
      Dir.chdir(tools_path)
      sh "curl -L https://github.com/appPlant/ff/releases/download/#{ENV["FF_VER"]}/ff-#{ENV["FF_VER"]}-x86_64-w64-mingw32.tgz  | tar xz"
    elsif OS.bits == 32
      sh"GOOS=windows GOARCH=386 go build /go/src/goo.go"
      Dir.chdir(tools_path)
      sh "curl -L https://github.com/appPlant/ff/releases/download/#{ENV["FF_VER"]}/ff-#{ENV["FF_VER"]}-i686-w64-mingw32.tgz  | tar xz"
    end
  end
end

namespace :test do
  desc "run mruby & unit tests"
  # only build mtest for host
  task :mtest => :compile do
    # in order to get mruby/test/t/synatx.rb __FILE__ to pass,
    # we need to make sure the tests are built relative from mruby_root
    MRuby.each_target do |target|
      # only run unit tests here
      target.enable_bintest = false
      run_test if target.test_enabled?
    end
  end

  def clean_env(envs)
    old_env = {}
    envs.each do |key|
      old_env[key] = ENV[key]
      ENV[key] = nil
    end
    yield
    envs.each do |key|
      ENV[key] = old_env[key]
    end
  end

  desc "run integration tests"
  task :bintest => :compile do
    ruby "/go/bintest/goo.rb"
  end
end

desc "run all tests"
Rake::Task['test'].clear
task :test => ["test:mtest", "test:bintest"]

desc "cleanup"
task :clean do
  sh "rake deep_clean"
end
=begin
desc "generate a release tarball"
task :release => :compile do
  require 'tmpdir'

  Dir.chdir(mruby_root) do
    # since we're in the mruby/
    release_dir  = "releases/v#{APP_VERSION}"
    release_path = Dir.pwd + "/../#{release_dir}"
    app_name     = "#{APP_NAME}-#{APP_VERSION}"
    FileUtils.mkdir_p(release_path)

    Dir.mktmpdir do |tmp_dir|
      Dir.chdir(tmp_dir) do
        MRuby.each_target do |target|
          next if name == "host"

          arch = name
          bin  = "#{build_dir}/bin/#{exefile(APP_NAME)}"
          FileUtils.mkdir_p(name)
          FileUtils.cp(bin, name)

          Dir.chdir(arch) do
            arch_release = "#{app_name}-#{arch}"
            puts "Writing #{release_dir}/#{arch_release}.tgz  | tar xz"
            `tar czf #{release_path}/#{arch_release}.tgz  | tar xz *`
          end
        end

        puts "Writing #{release_dir}/#{app_name}.tgz  | tar xz"
        `tar czf #{release_path}/#{app_name}.tgz  | tar xz *`
      end
    end
  end
end

namespace :local do
  desc "show version"
  task :version do
    puts APP_VERSION
  end
end

def is_in_a_docker_container?
  `grep -q docker /proc/self/cgroup`
  $?.success?
end

Rake.application.tasks.each do |task|
  next if ENV["MRUBY_CLI_LOCAL"]
  unless task.name.start_with?("local:")
    # Inspired by rake-hooks
    # https://github.com/guillermo/rake-hooks
    old_task = Rake.application.instance_variable_get('@tasks').delete(task.name)
    desc old_task.full_comment
    task old_task.name => old_task.prerequisites do
      abort("Not running in docker, you should type \"docker-compose run <task>\".") \
        unless is_in_a_docker_container?
      old_task.invoke
    end
  end
end
=end

