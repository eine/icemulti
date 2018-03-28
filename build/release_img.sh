#!/bin/sh

mkdir release
cd release
cp ../dtd-app.tgz ./
cp ../build/dockerfiles/prod ./Dockerfile
echo "ADD icemulti-app.tgz /icemulti" >> Dockerfile
docker build -t elide/icemulti:app .
cd ..
rm -rf release
