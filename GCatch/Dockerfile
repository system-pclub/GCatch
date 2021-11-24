FROM golang:1.16.4

RUN git clone https://github.com/Z3Prover/z3 /repos/z3
WORKDIR /repos/z3
RUN python scripts/mk_make.py \ 
  && cd build \
  && make \
  && make install

COPY . /go/src/github.com/system-pclub/GCatch/GCatch
RUN /go/src/github.com/system-pclub/GCatch/GCatch/install.sh

ENTRYPOINT [ "bash" ]