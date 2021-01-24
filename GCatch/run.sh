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
    echo "Z3 is not installed in $Z3. Please run installZ3.sh with sudo or checkout https://github.com/Z3Prover/z3 to install Z3"
    exit 1
  fi
  cd ../../../../..
  GCATCH=`pwd`
  GCATCH=$GCATCH/bin/GCatch
  if test -f "$GCATCH"; then
    echo "GCatch exists"
  else
    echo "GCatch is not installed in $GCATCH. Please run install.sh to install GCatch"
    exit 2
  fi
  echo "Step 1: setting GOPATH of the checked grpc"
  export GOPATH=$CURDIR/testdata/grpc-buggy
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
  echo "Step 2: running GCatch on a buggy version of grpc in testdata"
  echo "Note: all bugs reported below should be real BMOC bugs"
  echo "$GCATCH -path=$GOPATH/src/google.golang.org/grpc -include=google.golang.org/grpc -checker=BMOC -r"
  $GCATCH -path=$GOPATH/src/google.golang.org/grpc -include=google.golang.org/grpc -checker=BMOC -r

fi