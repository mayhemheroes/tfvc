# FROM fuzzers/go-fuzz:1.2.0
FROM golang:1.19 as builder

COPY . /tfvc
WORKDIR /tfvc/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

ENV GOFLAGS -mod=mod

RUN go-fuzz-build -libfuzzer -o fuzz.a && \
    clang -fsanitize=fuzzer fuzz.a -o fuzz.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /tfvc/mayhem/fuzz.libfuzzer /