#!/bin/bash -xe
cd "$(dirname "$0")"

docker build -t gcatch:local .

docker run -it --rm \
-v $(pwd)/tmp/playground:/playground \
-v $(pwd)/tmp/pkgmod:/go/pkg/mod \
gcatch:local $@