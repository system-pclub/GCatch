#!/bin/bash
mkdir bin
go build -o ./bin/dispatcher ../dispatcher/cmd/staticchecker
go build -o ./bin/gl1_patch ../gl-1-patcher/
go build -o ./bin/gl2_patch ../gl-2-patcher/
go build -o ./bin/gl3_patch ../gl-3-patcher/
export GOPATH=`pwd`
echo "running go get go.etcd.io/etcd"
go get go.etcd.io/etcd
cd src/go.etcd.io/etcd/
git reset --hard 6991f619f21375aac310aba4b5934105b2fbff9e
BASEPATH=`pwd`
cd $GOPATH

echo "dispatch and patch a GL-1 bug in pkg/transport/listener_test.go:"
echo original buggy code piece:
sed -n '149,182p' ${BASEPATH}/pkg/transport/listener_test.go
echo
echo dispatch result:
./bin/dispatcher -buggyfilepath=${BASEPATH}/pkg/transport/listener_test.go -path=${BASEPATH}/pkg/transport -makelineno=149 -oplineno=152 -include=go.etcd.io/etcd 2>&1 | grep DISPATCH
echo patch result:
./bin/gl1_patch ${BASEPATH}/pkg/transport/listener_test.go 149  | sed -n '132,166p'


echo "dispatch and patch a GL-2 bug in pkg/transport/timeout_dialer_test.go:"
echo original code:
sed -n '23,86p' ${BASEPATH}/pkg/transport/timeout_dialer_test.go
echo
echo dispatch result:
./bin/dispatcher -buggyfilepath=${BASEPATH}/pkg/transport/timeout_dialer_test.go -path=${BASEPATH}/pkg/transport -makelineno=24 -oplineno=102 -include=go.etcd.io/etcd 2>&1 | grep PATCH
echo patch result:
./bin/gl2_patch ${BASEPATH}/pkg/transport/timeout_dialer_test.go 26 85 | sed -n '9,72p'

echo "dispatch and patch a GL-3 bug in pkg/transport/timeout_dialer_test.go:"
echo original code:
sed -n '23,86p' ${BASEPATH}/pkg/transport/timeout_dialer_test.go
echo
echo dispatch result:
./bin/dispatcher -buggyfilepath=${BASEPATH}/pkg/transport/timeout_dialer_test.go -path=${BASEPATH}/pkg/transport -makelineno=45 -oplineno=48 -include=go.etcd.io/etcd 2>&1 | grep PATCH
echo patch result:
./bin/gl3_patch ${BASEPATH}/pkg/transport/timeout_dialer_test.go 45 48 72 | sed -n '25,88p'

rm -rf bin
echo removed bin dir