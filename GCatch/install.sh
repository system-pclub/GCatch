#!/usr/bin/env bash
#
# Install GCatch
# Before running this script, run installZ3.sh to install z3 from sources
# GCatch assumes that it has been installed under GOPATH and sets GOPATH accordingly
# GCatch currently only supports GOPATH instead of go.mod due to legacy issues
# GCatch turns off the usage of go module
# GCatch will be installed under GOPATH/bin/

echo "Make sure you have run installZ3.sh"

# cd script directory
CURDIR="$(dirname "$(realpath "$0")")"
cd "$CURDIR" || exit 1

# check if under GOPATH
REPATH='/src/github.com/system-pclub/GCatch/GCatch'
if [[ "$CURDIR" != *"$REPATH"* ]]
then
  echo "Please make sure the current directory is /SOME/PATH/src/github.com/system-pclub/GCatch/GCatch"
  echo "Current directory: $CURDIR"
  exit 1
else
  echo "Step 1: setting GOPATH to install GCatch"
  GOPATH=$(cd ../../../../.. || exit 1; pwd)
  export GOPATH
  echo "GOPATH is set to $GOPATH"
  echo "Step 2: installing GCatch"
  cd "$CURDIR"/cmd/GCatch || exit 1
  # turn off go mod before installation
  export GO111MODULE=off
  go install
  echo "GCatch is installed in $GOPATH/bin/GCatch"
fi
