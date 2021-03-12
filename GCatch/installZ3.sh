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
    echo "Installing Z3 to a default position, normally $Z3"
    echo "Could fail if ran without sudo"
    cd ./tools/z3
    python scripts/mk_make.py
    cd build
    make
    sudo make install
  fi
fi