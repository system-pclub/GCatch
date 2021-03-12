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
  echo "Step 1: setting GOPATH to install GCatch"
  cd ../../../../..
  export GOPATH=`pwd`
  echo "GOPATH is set to $GOPATH"
  echo "Step 2: installing GCatch"
  cd $CURDIR/cmd/GCatch
  go install
  echo "GCatch is installed in $GOPATH/bin/GCatch"
fi