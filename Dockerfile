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

FROM golang:alpine
MAINTAINER Sebastian Katzer "katzer@appplant.de"

ENV APP_HOME /root/code
RUN mkdir $APP_HOME
WORKDIR $APP_HOME

COPY scripts/install.sh .
RUN sh install.sh
RUN gem install rake os test-unit --no-ri --no-rdoc
RUN gem update --system --no-ri --no-rdoc
RUN git config --global http.https://gopkg.in.followRedirects true
RUN go get gopkg.in/hypersleep/easyssh.v0
RUN go get github.com/Sirupsen/logrus
RUN go get github.com/olekukonko/tablewriter
RUN git -C $GOPATH/src/golang.org/x/crypto reset --hard abc5fa7ad02123a41f02bf1391c9760f7586e608
RUN apk add tree

COPY scripts/init.sh /etc/profile.d
