FROM golang:alpine

# ENV GOBIN /go/bin
# ENV TOOLS_PATH /go/tools
# ENV FF_VER 0.0.1
#ENV IPS_ORBIT_FILE /go/bintest/testtools/test.json
# ENV PATH $PATH:/go/tools
ENV APP_VERSION 0.0.1
ENV APP_NAME goo
ENV BUILD_VERSION 0.0.1



RUN apk update
RUN apk add ruby
# RUN apk add ruby-rdoc
# RUN apk add ruby-dev
# RUN apk add ruby-irb
RUN apk add ruby-rake
# RUN apk add bash
#RUN apk add bash-doc
# RUN apk add bash-completion
RUN apk add openssh
RUN gem install os --no-ri --no-rdoc
RUN apk add wget
RUN apk --no-cache add ca-certificates
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.23-r3/glibc-2.23-r3.apk
RUN apk add glibc-2.23-r3.apk
RUN /usr/bin/ssh-keygen -A
RUN /usr/sbin/sshd
RUN mkdir /root/.ssh
RUN ssh-keygen -q -f /root/.ssh/id_rsa -N ""
RUN mv /root/.ssh/id_rsa.pub /root/.ssh/authorized_keys

ENV GOROOT /usr/local/go
ENV GOPATH /usr/local/go/packages
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

RUN apk add git
RUN go get golang.org/x/crypto/ssh




