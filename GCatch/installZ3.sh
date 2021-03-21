#!/usr/bin/env bash
# Install z3 from sources
# GCatch requires z3 WITH SOURCES

installZ3() {
  echo "Installing z3 to a default position, normally $Z3"
  echo "Could fail if run without sudo"
  cd ./tools/z3 || exit
  python scripts/mk_make.py
  cd build || exit
  make
  sudo make install
}

# cd script directory
cd "$(dirname "$(realpath "$0")")" || exit;
Z3=/usr/local/bin/z3
if test -f "$Z3"; then
  read -p 'GCatch requires z3 WITH SOURCES. Reinstall z3 from sources? [Y/n] ' -r yn
  case $yn in
    [Yy]* ) installZ3;;
    [Nn]* ) exit;;
    * ) echo "Please answer Y/N";;
  esac
else
  installZ3
fi
