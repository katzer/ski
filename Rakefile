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
APP_VERSION=ENV["APP_VERSION"] || "6.6.6"
bin_path="#{APP_ROOT}/bin"
src_path="#{APP_ROOT}/src"
tools_path="#{ENV["TOOLS_PATH"]}"

desc "compile binary"
task :compile do
  FileUtils.rm_r "#{APP_ROOT}/tools" unless !Dir.exists?("#{APP_ROOT}/tools")
  Dir.mkdir("#{APP_ROOT}/tools")
  FileUtils.rm_r bin_path unless !Dir.exists?(bin_path)
  Dir.mkdir(bin_path)
  Dir.chdir(bin_path)
  Dir.mkdir("Linux64")
  Dir.mkdir("Linuxi686")
  Dir.mkdir("Win64")
  Dir.mkdir("Wini686")
  Dir.mkdir("Mac64")
  Dir.mkdir("Maci386")
  Dir.chdir("#{bin_path}/Linux64")
  puts "Linux64 #{system("GOOS=linux GOARCH=amd64 go build #{src_path}/goo.go")}"
  Dir.chdir("#{bin_path}/Linuxi686")
  puts "Linuxi686 #{system("GOOS=linux GOARCH=386 go build #{src_path}/goo.go")}"
  Dir.chdir("#{bin_path}/Win64")
  puts "Win64 #{system("GOOS=windows GOARCH=amd64 go build #{src_path}/goo.go")}"
  Dir.chdir("#{bin_path}/Wini686")
  puts "Wini686 #{system("GOOS=windows GOARCH=386 go build #{src_path}/goo.go")}"
  Dir.chdir("#{bin_path}/Mac64")
  puts "Mac64 #{system("GOOS=darwin GOARCH=amd64 go build #{src_path}/goo.go")}"
  Dir.chdir("#{bin_path}/Maci386")
  puts "Maci386 #{system("GOOS=darwin GOARCH=386 go build #{src_path}/goo.go")}"

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
    os = ""
    if OS.linux?
      if OS.bits == 64
        os = "Linux64"
      elsif OS.bits == 32
        os = "Linuxi686"
      end
    elsif OS.mac?
      if OS.bits == 64
        os = "Mac64"
      elsif OS.bits == 32
        os = "Maci386"
      end
    elsif OS.windows?
      if OS.bits == 64
        os = "Win64"
      elsif OS.bits == 32
        os = "Wini686"
      end
    end
    bin_path = File.join(File.dirname(__FILE__), "bin/#{os}/goo")
    ruby "#{APP_ROOT}/bintest/goo.rb #{bin_path}"
  end
end

desc "run all tests"
Rake::Task['test'].clear
task :test => ["test:mtest", "test:bintest"]

desc "cleanup"
task :clean do
  sh "rake deep_clean"
end

desc "generate a release tarball"
task :release => :compile do

  Dir.chdir(APP_ROOT)
  release_dir  = "releases/v#{APP_VERSION}"
  release_path = Dir.pwd + "/#{release_dir}"
  app_name     = "#{APP_NAME}-#{APP_VERSION}"
  FileUtils.mkdir_p(release_path)

  #Linux64
  arch = "x86_64-pc-linux-gnu"
  arch_release = "#{app_name}-#{arch}"
  Dir.chdir("#{bin_path}/Linux64")
  puts "#{bin_path}/Linux64"
  puts "Writing #{release_dir}/#{arch_release}.tgz "
  `tar czf #{release_path}/#{arch_release}.tgz *`
  #Linuxi686
  arch = "i686-pc-linux-gnu"
  arch_release = "#{app_name}-#{arch}"
  Dir.chdir("#{bin_path}/Linuxi686")
  puts "Writing #{release_dir}/#{arch_release}.tgz "
  `tar czf #{release_path}/#{arch_release}.tgz *`
  #Mac64
  arch = "x86_64-apple-darwin14"
  arch_release = "#{app_name}-#{arch}"
  Dir.chdir("#{bin_path}/Mac64")
  puts "Writing #{release_dir}/#{arch_release}.tgz "
  `tar czf #{release_path}/#{arch_release}.tgz *`
  #Maci386
  arch = "i386-apple-darwin14"
  arch_release = "#{app_name}-#{arch}"
  Dir.chdir("#{bin_path}/Maci386")
  puts "Writing #{release_dir}/#{arch_release}.tgz "
  `tar czf #{release_path}/#{arch_release}.tgz *`
  #Win64
  arch = "x86_64-w64-mingw32"
  arch_release = "#{app_name}-#{arch}"
  Dir.chdir("#{bin_path}/Win64")
  puts "Writing #{release_dir}/#{arch_release}.tgz "
  `tar czf #{release_path}/#{arch_release}.tgz *`
  #Wini686
  arch = "i686-w64-mingw32"
  arch_release = "#{app_name}-#{arch}"
  Dir.chdir("#{bin_path}/Wini686")
  puts "Writing #{release_dir}/#{arch_release}.tgz "
  `tar czf #{release_path}/#{arch_release}.tgz *`

  #source
  Dir.chdir(src_path)
  puts "Writing #{release_dir}/#{app_name}.tgz "
  `tar czf #{release_path}/#{app_name}.tgz *`

end
=begin
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

