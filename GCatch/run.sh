#!/usr/bin/env bash
#
# Run GCatch to check the buggy grpc
# Before running this script, run installZ3.sh and then install.sh
# GCatch currently only supports GOPATH instead of go.mod due to legacy issues
# GCatch turns off the usage of go module
# The project to be checked must be located in the corresponding GOPATH

echo "Make sure you have run installZ3.sh and then install.sh"
echo "This script takes two parameters:"
echo "\tThe first parameter should be the GOPATH of the target program you want to verify"
echo "\tThe second parameter should be the relative path (what is after GOPATH/src/) of the target program"

# check if z3 installed
if ! command -v z3 >/dev/null; then
  echo "Cannot detect z3. Run installZ3.sh to install z3 from sources"
fi

# check if z3 header installed
if ! echo '#include <z3.h>' | gcc -H -fsyntax-only -E - 1>/dev/null 2>&1; then
  echo "Cannot detect <z3.h>. Run installZ3.sh to install z3 from sources"
fi

# cd script directory
CURDIR="$(dirname "$(realpath "$0")")"
cd "$CURDIR" || exit 1

GCATCH="$(cd ../../../../.. || exit 1; pwd)/bin/GCatch"

# check if GCatch installed
if ! test -f "$GCATCH"; then
  echo "GCatch is not installed. Run install.sh to install it under $GCATCH"
  exit 1
fi

# turn off go mod before checking
export GO111MODULE=off

echo "Step 1: setting GOPATH"
export GOPATH=$1
echo "GOPATH is set to $GOPATH"
echo ""
echo "Step 2: running GCatch on the input program"
echo "$GCATCH -path=$GOPATH/src/$2 -include=$2 -checker=BMOC -compile-error"
$GCATCH -path="$GOPATH"/src/$2 -include=$2 -checker=BMOC -compile-error
