#!/bin/sh

cd $(dirname $0)
if [ -f ./utils.sh ]; then . ./utils.sh; fi

docker build -t elide/icemulti:dev - <<EOF
FROM golang:alpine
RUN apk add -U --no-cache git yarn graphviz curl \
 && curl -L \$( \
      curl -s https://api.github.com/repos/gohugoio/hugo/releases/latest \
    | grep "browser_download_url.*Linux-64bit.*tar.gz" \
    | cut -d : -f 2,3 \
    | tr -d \" \
  ) \
  | tar -xvz hugo -C /usr/bin/
EOF

#travis_start alpine-factory
#  docker build -t elide/alpine:factory --target alpine-factory -f dockerfiles/alpine .
#travis_finish alpine-factory
#
#travis_start alpine-build
#  mkdir -p tmp-pkgs
#  cd tmp-pkgs
#  docker run --rm -tv /$(pwd)://work -w //work elide/alpine:factory sh -c "$(cat ../utils.sh ../build.sh)"
#travis_finish alpine-build

#docker build -t elide/alpine --target alpine-build -f dockerfiles/alpine .
#docker build -t elide/alpine --target alpine-release -f dockerfiles/alpine .
