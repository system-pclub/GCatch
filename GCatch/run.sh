#!/usr/bin/env bash
#
# Run GCatch to check the buggy grpc
# Before running this script, run installZ3.sh and then install.sh
# GCatch currently only supports GOPATH instead of go.mod due to legacy issues
# GCatch turns off the usage of go module
# The project to be checked must be located in the corresponding GOPATH
echo "Running experiments for ECOOP submission"

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

echo GCatch path: $GCATCH

# turn off go mod before checking
export GO111MODULE=off

echo "Step 1: setting GOPATH"
export GOPATH=$CURDIR/testdata/ecoop
echo "GOPATH is set to $GOPATH"
echo ""
echo "Step 2: running GCatch on 7 input programs"

listVar="figure1 figure2 figure2_translate figure12_1 figure12_2 figure12_3 figure13_1 figure13_2"
# listVar="figure2 figure2_translate"
for i in $listVar; do
  echo "========================================================"
  echo "===============Running GCatch on $i============="
    echo "========================================================"
    echo "$GCATCH -path=$GOPATH/src/$i -include=$i -checker=BMOC -compile-error"
    $GCATCH -path="$GOPATH"/src/$i -include=$i -checker=BMOC -compile-error
    echo ""
    echo ""
done

echo "Step 1: setting GOPATH"
export GOPATH=$CURDIR/testdata/gobench
echo "GOPATH is set to $GOPATH"
echo ""
echo "Step 2: running GCatch on 7 input programs"

listVar="etcd/7902 etcd/6873 etcd/7492 etcd/7443 istio/16224 moby/28462 kubernetes/6632 kubernetes/26980 kubernetes/10182 kubernetes/1321 serving/2137 grpc/1353 grpc/1460"
for i in $listVar; do
  echo "========================================================"
  echo "===============Running GCatch on $i============="
    echo "========================================================"
    echo "$GCATCH -path=$GOPATH/src/$i -include=$i -checker=BMOC -compile-error"
    $GCATCH -path="$GOPATH"/src/$i -include=$i -checker=BMOC -compile-error
    echo ""
    echo ""
done
 