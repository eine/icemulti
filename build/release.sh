#!/bin/sh

if [ -d "release" ]; then
  rm -rf release
fi
mkdir release

cd app
if [ -d "dist" ]; then
  rm -rf dist
fi
yarn build -v
mv dist ../release/public

cd ../api
go build -o ../release/icemulti

cd ..

tar -zcvf icemulti-app.tgz -C release .

rm -rf release
