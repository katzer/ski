FROM golang:alpine

ENV GOBIN /go/bin
ENV TOOLS_PATH /go/tools
ENV FF_VER 0.0.1
ENV IPS_ORBIT_FILE /go/bintest/test.json
ENV PATH $PATH:/go/tools

RUN apk update
RUN apk add curl
RUN apk add ruby
RUN apk add ruby-rdoc
RUN apk add ruby-dev
RUN apk add ruby-irb
RUN apk add ruby-rake
RUN apk add bash bash-doc bash-completion
RUN gem install os
RUN apk add wget
RUN apk --no-cache add ca-certificates
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.23-r3/glibc-2.23-r3.apk
RUN apk add glibc-2.23-r3.apk

