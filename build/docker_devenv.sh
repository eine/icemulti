#!/bin/sh

set -e # Exit with nonzero exit code if anything fails

cd $(dirname $0)

HOST_DIR=$(dirname $(pwd))
REPO_BIND="/src"
WORK_DIR="/go/src/github.com/1138-4EB"
PORT="5555:8080"

echo "[icemulti devenv] Launch elide/icemulti:dev docker container"

echo "Host dir      | $HOST_DIR"
echo "Work (go) dir | $WORK_DIR"
echo "Port map/bind | $PORT"

if [ "$(command -v dos2unix)" != "" ]; then
  dos2unix init_devenv.sh
  dos2unix release.sh
fi

$(command -v winpty) docker \
  run --rm -it \
  -v /"${HOST_DIR}":/"$REPO_BIND" \
  -e REPO_BIND=/"$REPO_BIND" \
  -p "$PORT" \
  -w /"$WORK_DIR" \
  elide/icemulti:dev sh -c "$(cat init_devenv.sh); $1"
