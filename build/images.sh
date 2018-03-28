#!/bin/sh

cd $(dirname $0)
if [ -f ./travis-utils.sh ]; then . ./travis-utils.sh; fi

fold_start alpine-factory
  docker build -t elide/alpine:factory --target alpine-factory -f dockerfiles/alpine .
fold_end alpine-factory

fold_start alpine-build
  mkdir -p tmp-pkgs
  cd tmp-pkgs
  docker run --rm -tv /$(pwd)://work -w //work elide/alpine:factory sh -c "$(cat ../travis-utils.sh ../build.sh)"
fold_end alpine-build

#docker build -t elide/alpine --target alpine-build -f dockerfiles/alpine .
#docker build -t elide/alpine --target alpine-release -f dockerfiles/alpine .

#docker build -t elide/icemulti:dev - < build/dockerfile/icemulti-dev

# docker build -t elide/icemulti:dev - < dockerfile/icemulti-dev
