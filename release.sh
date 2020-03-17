#!/bin/sh

set -e

cd $(dirname "$0")

if [ -d "release" ]; then
  rm -rf release
fi
mkdir release

cd app
if [ -d "dist" ]; then
  rm -rf dist
fi
yarn
yarn build -v
mv dist ../release/public

cd ../api
CGO_ENABLED=0 go build -a -o ../release/icemulti

cd ..

tar -zcvf icemulti-app.tgz -C release .

rm -rf release
