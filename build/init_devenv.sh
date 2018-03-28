#!/bin/sh

if [ "$REPO_BIND" != "" ]; then
  echo "[icemulti] symlink $REPO_BIND and work_dir"
  ln -s "$REPO_BIND" dtd
fi

echo "[icemulti] go get icemulti/api dependencies"

cd dtd/api
dep ensure

echo "[icemulti] yarn install icemulti/app dependencies"

cd ../app
yarn install

echo "[icemulti] go get icemulti/cli dependencies"

cd ../cli
dep ensure

echo "[icemulti] go get icemulti/lib dependencies"

cd ../lib
dep ensure

cd ..

echo "[icemulti] Happy hacking!"
