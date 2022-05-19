#!/usr/bin/env bash
#
# Run GCatch to check two toy programs with non-blocking bugs in testdata/toyprogram/src/*Close
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

echo "-----Step 1: setting GOPATH of the toyprogram"
export GOPATH=$CURDIR/testdata/toyprogram
echo "GOPATH is set to $GOPATH"
echo ""
echo "Description of flags of GCatch:"
echo "Required Flag: -path=Full path of the application to be checked"
echo "Required Flag: -include=Relative path (what's after /src/) of the application to be checked"
echo "Required Flag: -checker=The checkers you want to run, divided by \":\".    Default value:BMOC"
echo "Optional Flag: -r    Whether all children packages should also be checked recursively"
echo "Optional Flag: -compile-error    Whether compilation errors should be printed, if there are any"
echo "Optional Flag: -vendor=Packages that will be ignored, divided by \":\".    Default value:vendor"
echo ""

echo "-----Step 2.1: running GCatch on testdata/toyprogram/src/doubleClose"
echo "GO111MODULE=off $GCATCH -path="$GOPATH"/src/doubleClose -include=doubleClose -checker=BMOC:NBMOC -r"
GO111MODULE=off $GCATCH -path="$GOPATH"/src/doubleClose -include=doubleClose -checker=BMOC:NBMOC -r
echo ""
echo "-----Step 2.2: running GCatch on testdata/toyprogram/src/sendAfterClose"
echo "GO111MODULE=off $GCATCH -path="$GOPATH"/src/sendAfterClose -include=sendAfterClose -checker=BMOC:NBMOC -r"
GO111MODULE=off $GCATCH -path="$GOPATH"/src/sendAfterClose -include=sendAfterClose -checker=BMOC:NBMOC -r
