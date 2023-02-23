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

# echo "Step 1: setting GOPATH"
# export GOPATH=$CURDIR/testdata/ecoop
# echo "GOPATH is set to $GOPATH"
# echo ""
# echo "Step 2: running GCatch on 7 input programs"

# listVar="figure1 figure2 figure2_translate figure12_1 figure12_2 figure12_3 figure13_1 figure13_2"
# listVar="figure2 figure2_translate"
# listVar="feb_6"
# for i in $listVar; do
#   echo "========================================================"
#   echo "===============Running GCatch on $i============="
#     echo "========================================================"
#     echo "$GCATCH -path=$GOPATH/src/$i -include=$i -checker=BMOC -compile-error"
#     $GCATCH -path="$GOPATH"/src/$i -include=$i -checker=BMOC -compile-error
#     echo ""
#     echo ""
# done


# echo "Step 1: setting GOPATH"
# export GOPATH=$CURDIR/testdata/gobench
# echo "GOPATH is set to $GOPATH"
# echo ""
# echo "Step 2: running GCatch on 7 input programs"

# listVar="etcd/7902 etcd/6873 etcd/7492 etcd/7443 istio/16224 moby/28462 kubernetes/6632 kubernetes/26980 kubernetes/10182 kubernetes/1321 serving/2137 grpc/1353 grpc/1460"
# listVar="grpc/1687 serving/5865 serving/3068"
# listVar="grpc/2371 etcd/3077 istio/8967"
# for i in $listVar; do
#   echo "========================================================"
#   echo "===============Running GCatch on $i============="
#     echo "========================================================"
#     echo "$GCATCH -path=$GOPATH/src/$i -include=$i -checker=BMOC -compile-error"
#     $GCATCH -path="$GOPATH"/src/$i -include=$i -checker=BMOC -compile-error
#     echo ""
#     echo ""
# done
# exit


echo "Step 1: setting GOPATH"
export GOPATH=$CURDIR/testdata/gobench-full
echo "GOPATH is set to $GOPATH"
echo ""
echo "Step 2: running GCatch on 7 input programs"

# listVar="etcd/7902 etcd/6873 etcd/7492 etcd/7443 istio/16224 moby/28462 kubernetes/6632 kubernetes/26980 kubernetes/10182 kubernetes/1321 serving/2137 grpc/1353 grpc/1460"
listVarNonBlocking=$(<<EOF
nonblocking/etcd/9446
nonblocking/etcd/8194
nonblocking/etcd/4876
nonblocking/etcd/3077
nonblocking/istio/8967
nonblocking/istio/8144
nonblocking/istio/16742
nonblocking/istio/8214
nonblocking/moby/27037
nonblocking/moby/18412
nonblocking/moby/22941
nonblocking/kubernetes/88331
nonblocking/kubernetes/82550
nonblocking/kubernetes/81148
nonblocking/kubernetes/70892
nonblocking/kubernetes/89164
nonblocking/kubernetes/77796
nonblocking/kubernetes/49404
nonblocking/kubernetes/81091
nonblocking/kubernetes/82239
nonblocking/kubernetes/79631
nonblocking/kubernetes/13058
nonblocking/kubernetes/80284
nonblocking/serving/6171
nonblocking/serving/6472
nonblocking/serving/5865
nonblocking/serving/3068
nonblocking/serving/4908
nonblocking/serving/3148
nonblocking/grpc/3090
nonblocking/grpc/2371
nonblocking/grpc/1687
nonblocking/grpc/1748
nonblocking/cockroach/35501
nonblocking/cockroach/4407
EOF
)

listVar="grpc/1687 serving/5865 serving/3068"
listVar='blocking/hugo/3251
blocking/hugo/5379
blocking/etcd/7902
blocking/etcd/5509
blocking/etcd/6873
blocking/etcd/6708
blocking/etcd/6857
blocking/etcd/7492
blocking/etcd/10492
blocking/etcd/7443
blocking/istio/18454
blocking/istio/17860
blocking/istio/16224
blocking/moby/33293
blocking/moby/25384
blocking/moby/27782
blocking/moby/28462
blocking/moby/36114
blocking/moby/17176
blocking/moby/33781
blocking/moby/29733
blocking/moby/7559
blocking/moby/30408
blocking/moby/21233
blocking/moby/4951
blocking/moby/4395
blocking/kubernetes/11298
blocking/kubernetes/6632
blocking/kubernetes/5316
blocking/kubernetes/13135
blocking/kubernetes/62464
blocking/kubernetes/38669
blocking/kubernetes/26980
blocking/kubernetes/58107
blocking/kubernetes/25331
blocking/kubernetes/10182
blocking/kubernetes/1321
blocking/kubernetes/30872
blocking/kubernetes/70277
blocking/syncthing/4829
blocking/syncthing/5795
blocking/serving/2137
blocking/grpc/795
blocking/grpc/1424
blocking/grpc/3017
blocking/grpc/1275
blocking/grpc/1353
blocking/grpc/660
blocking/grpc/862
blocking/grpc/1460
blocking/cockroach/13197
blocking/cockroach/35073
blocking/cockroach/1055
blocking/cockroach/10790
blocking/cockroach/1462
blocking/cockroach/6181
blocking/cockroach/10214
blocking/cockroach/584
blocking/cockroach/18101
blocking/cockroach/2448
blocking/cockroach/24808
blocking/cockroach/9935
blocking/cockroach/35931
blocking/cockroach/25456
blocking/cockroach/16167
blocking/cockroach/7504
blocking/cockroach/3710
blocking/cockroach/13755'

#listVar='blocking/hugo/5379'
echo listvar= $listVar
for i in $listVar; do
  echo "========================================================"
  echo "===============Running GCatch on $i============="
    echo "========================================================"
    echo "$GCATCH -path=$GOPATH/src/$i -include=$i -checker=BMOC -compile-error"
    $GCATCH -path="$GOPATH"/src/$i -include=$i -checker=unlock:double:conflict -compile-error || echo "GCatch exited with non-zero exit code."
    echo ""
    echo ""
done