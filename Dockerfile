FROM golang:1.16.0-buster

RUN apt-get update && apt-get install -y --no-install-recommends \
  git \
  curl \
  less \
  vim

# vscodeのgolang.goの依存ライブラリ
RUN go get github.com/uudashr/gopkgs \
  github.com/ramya-rao-a/go-outline  \
  github.com/cweill/gotests \
  github.com/fatih/gomodifytags \
  github.com/josharian/impl \
  github.com/haya14busa/goplay \
  github.com/go-delve/delve \
  github.com/go-delve/delve/cmd/dlv \
  golang.org/x/lint/golint \
  golang.org/x/tools/gopls

RUN mkdir -p /go/src/github.com/sainu/profile-api

WORKDIR /go/src/github.com/sainu/profile-api

ADD . .
