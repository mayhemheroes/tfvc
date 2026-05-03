FROM golang:1.25 AS builder

COPY . /tfvc
WORKDIR /tfvc/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN GOFLAGS=-mod=mod go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

ENV GOFLAGS=-mod=mod

RUN go-fuzz-build -libfuzzer -o fuzzLockFile.a &&     clang -fsanitize=fuzzer fuzzLockFile.a -o fuzzLockFile.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /tfvc/mayhem/fuzzLockFile.libfuzzer /
