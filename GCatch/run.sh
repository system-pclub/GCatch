CURDIR=`pwd`
REPATH='/src/github.com/system-pclub/GCatch/GCatch'
if [[ "$CURDIR" != *"$REPATH"* ]]
then
  echo "Please make sure the current directory is /SOME/PATH/src/github.com/system-pclub/GCatch/GCatch"
  echo "Current directory: $CURDIR"
  exit 0
else
  Z3=/usr/local/bin/z3
  if test -f "$Z3"; then
    echo "Z3 exists"
  else
    echo "Z3 is not installed in $Z3. Please check README.md or https://github.com/Z3Prover/z3 to install Z3"
    exit 1
  fi
  echo "Step 1: setting GOPATH to install GCatch"
  cd ../../../../..
  export GOPATH=`pwd`
  echo "GOPATH is set to $GOPATH"
  echo "Step 2: installing GCatch"
  cd $CURDIR/cmd/GCatch
  go install
  echo "GCatch is installed in $GOPATH/bin/GCatch"
  GCATCH=$GOPATH/bin/GCatch
  echo "Step 3: setting GOPATH of the checked grpc"
  export GOPATH=$CURDIR/testdata/grpc-buggy
  echo "GOPATH is set to $GOPATH"
  echo "Step 4: running GCatch on $CURDIR/testdata/grpc-buggy/src/google.golang.org/grpc"
  echo "Note: all bugs reported below should be real BMOC bugs"
  $GCATCH -path=$GOPATH/src/google.golang.org/grpc -include=google.golang.org/grpc -checker=BMOC -r
fi