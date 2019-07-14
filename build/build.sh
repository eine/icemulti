#!/bin/sh

cd $(dirname $0)
if [ -f utils.sh ]; then . utils.sh; fi

travis_start steveicarus/iverilog
  curl -L https://github.com/steveicarus/iverilog/archive/master.tar.gz | tar xzv
  cd iverilog-master
  sh autoconf.sh
  ./configure DESTDIR=$HOME/package-iverilog
  make -j$(nproc) DESTDIR=$HOME/package-iverilog
  make check

  cd ..
travis_finish steveicarus/iverilog

travis_start cliffordwolf/icestorm
  curl -L https://github.com/cliffordwolf/icestorm/archive/master.tar.gz | tar xvz
  cd icestorm-master
  make -j$(nproc) DESTDIR=$HOME/package-icestorm

  cd ..
travis_finish cliffordwolf/icestorm

travis_start cseed/arachne-pnr
  curl -L https://github.com/cseed/arachne-pnr/archive/master.tar.gz | tar xvz
  cd arachne-pnr-master
  make -j$(nproc) DESTDIR=$HOME/package-arachne-pnr

  cd ..
travis_finish cseed/arachne-pnr

travis_start cliffordwolf/yosys
  curl -L https://github.com/cliffordwolf/yosys/archive/master.tar.gz | tar xvz
  cd yosys-master
  make -j$(nproc) DESTDIR=$HOME/package-yosys
  make test

  cd ..
travis_finish cliffordwolf/yosys

travis_start gtkwave-code/gtkwave
#REPO=svn://svn.code.sf.net/p/gtkwave/code/gtkwave3
#svn checkout $REPO ./

  curl -L https://github.com/gtkwave-code/gtkwave/archive/master.tar.gz | tar xvz
  cd gtkwave-master
  ./configure --prefix=/usr/local --with-tk=/usr/lib
  make -j$(nproc) DESTDIR=$HOME/package-gtkwave
  make check

  cd ..
travis_finish gtkwave-code/gtkwave

#cd $HOME/package-$TOOL/usr/local/
#tar -cvf $HOME/package-$TOOL/$TOOL.tar *
#cd $HOME && rm -rf $HOME/package-$TOOL/usr
#tar -zcvf $HOME/$ELIDE_TAG-$TOOL-$TOOL_SHORT.tar.gz package-$TOOL/*
#rm -rf $TOOL package-$TOOL
