#!/bin/sh

set -e # Exit with nonzero exit code if anything fails

cd $(dirname $0)

HOST_DIR=$(dirname $(pwd))
PORT="5555:8080"

echo "[icemulti devenv] Launch elide/icemulti:dev docker container"

echo "Host dir      | $HOST_DIR"
echo "Port map/bind | $PORT"

if [ "$(command -v dos2unix)" != "" ]; then
  dos2unix release.sh
fi

$(command -v winpty) docker \
  run --rm -it \
  -p "$PORT" \
  -v /"${HOST_DIR}"://src \
  -w //src \
  elide/icemulti:dev sh -c "go get ./... ; cd app; yarn install; cd ..; $1"
