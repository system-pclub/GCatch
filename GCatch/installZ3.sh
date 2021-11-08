#!/usr/bin/env bash
#
# Install z3 from sources
# GCatch requires z3 WITH SOURCES
# You can also checkout https://github.com/Z3Prover/z3 to install Z3

installZ3() {
  echo "Installing z3 to a default position, normally $Z3"
  echo "Could fail if run without sudo"
  cd ./tools/z3 || exit 1
  python scripts/mk_make.py
  cd build || exit 1
  make
  sudo make install
}

# cd script directory
cd "$(dirname "$(realpath "$0")")" || exit 1

# check z3 bin
if ! Z3=$(command -v z3); then
  echo 'Cannot detect z3'
  installZ3
  exit
fi

# check z3 header
if echo '#include <z3.h>' | gcc -H -fsyntax-only -E - 1>/dev/null 2>&1; then
  read -p 'z3 and <z3.h> exist. Reinstall z3 from sources? [y/N] ' -r yn
  case $yn in
    [Yy]* ) installZ3;;
    * ) exit;;
  esac
else
  read -p 'GCatch requires z3 WITH SOURCES. Reinstall z3 from sources? [Y/n] ' -r yn
  case $yn in
    [Nn]* ) exit;;
    * ) installZ3;;
  esac
fi
