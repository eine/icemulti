#! /bin/sh

set -e

cd $(dirname $0)
if [ -f ./utils.sh ]; then . ./utils.sh; fi

echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin

for key in $FILTER; do
  for tag in `echo $(docker images elide/icemulti* | awk -F ' ' '{print $1 ":" $2}') | cut -d ' ' -f2-`; do
      if [ "$tag" = "REPOSITORY:TAG" ]; then break; fi
      id="`echo $tag | grep -oP 'ghdl/\K.*'`"
      travis_start "$id"
      printf "$ANSI_YELLOW[DOCKER push] ${tag}$ANSI_NOCOLOR\n"
      docker push $tag
      travis_finish "$id"
  done
done
docker logout
