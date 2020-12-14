#!/bin/bash
mkdir bin
go build -o ./bin/dispatcher ../dispatcher/cmd/staticchecker
go build -o ./bin/gl1_patch ../gl-1-patcher/
go build -o ./bin/gl2_patch ../gl-2-patcher/
go build -o ./bin/gl3_patch ../gl-3-patcher/
export GOPATH=`pwd`
echo "dispatch and patch the example for GL-1 bugs:"
echo original code:
cat $GOPATH/src/gl1/gl_1.go
echo
echo dispatch result:
./bin/dispatcher -buggyfilepath=$GOPATH/src/gl1/gl_1.go -path=$GOPATH -makelineno=9 -oplineno=12 -include=$GOPATH 2>&1 | grep DISPATCH
echo patch result:
./bin/gl1_patch `pwd`/src/gl1/gl_1.go 9


echo "dispatch and patch the example for GL-2 bugs:"
echo original code:
cat $GOPATH/src/gl2/gl_2.go
echo
echo dispatch result:
./bin/dispatcher -buggyfilepath=$GOPATH/src/gl2/gl_2.go -path=$GOPATH -makelineno=9 -oplineno=12 -include=$GOPATH 2>&1  | grep PATCH
echo patch result:
./bin/gl2_patch `pwd`/src/gl2/gl_2.go 10 18


echo "dispatch and patch the example for GL-3 bugs:"
echo original code:
cat $GOPATH/src/gl3/gl_3.go
echo dispatch result:
./bin/dispatcher -buggyfilepath=$GOPATH/src/gl3/gl_3.go -path=$GOPATH -makelineno=9 -oplineno=14 -include=$GOPATH 2>&1  | grep PATCH
echo patch result:
./bin/gl3_patch `pwd`/src/gl3/gl_3.go 10 14


rm -rf bin
echo removed bin dir