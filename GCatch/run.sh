#!/usr/bin/env bash
#
# Run GCatch to check the buggy grpc
# Before running this script, run installZ3.sh and then install.sh
# GCatch currently only supports GOPATH instead of go.mod due to legacy issues
# GCatch turns off the usage of go module
# The project to be checked must be located in the corresponding GOPATH

echo "Make sure you have run installZ3.sh and then install.sh"

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
echo "Step 1: setting target"
export TARGET=$2
echo "GOPATH is set to $TARGET"
echo ""
echo "Description of flags of GCatch:"
echo "Required Flag: -path=Full path of the application to be checked"
echo "Required Flag: -include=Relative path (what's after /src/) of the application to be checked"
echo "Required Flag: -checker=The checkers you want to run, divided by \":\".    Default value:BMOC"
echo "Optional Flag: -r    Whether all children packages should also be checked recursively"
echo "Optional Flag: -compile-error    Whether compilation errors should be printed, if there are any"
echo "Optional Flag: -vendor=Packages that will be ignored, divided by \":\".    Default value:vendor"
echo ""
echo "Step 2: running GCatch on a buggy version of grpc in testdata"
echo "Note: all bugs reported below should be real BMOC bugs"
echo "GO111MODULE=off $GCATCH -path=$GOPATH/src/google.golang.org/grpc -include=google.golang.org/grpc -checker=BMOC -r"
GO111MODULE=off $GCATCH -path="$1/src/$2" -include="$2" -checker=BMOC -r
