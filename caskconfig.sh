#!/bin/sh
brew install go
brew install ruby
brew install wget
gem install os
export GOROOT=/usr/local/go
export GOPATH=/usr/local/go/packages
go get gopkg.in/hypersleep/easyssh.v0
