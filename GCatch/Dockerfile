FROM golang:1.16.4

RUN git clone https://github.com/Z3Prover/z3 /repos/z3
WORKDIR /repos/z3
RUN python scripts/mk_make.py \ 
  && cd build \
  && make \
  && make install

COPY . /go/src/github.com/system-pclub/GCatch/GCatch


RUN GO111MODULE=off go get golang.org/x/xerrors \
&& GO111MODULE=off go get golang.org/x/mod/semver \
&& GO111MODULE=off go get golang.org/x/sys/execabs

RUN /go/src/github.com/system-pclub/GCatch/GCatch/install.sh

WORKDIR /playground
ENTRYPOINT [ "bash" ]