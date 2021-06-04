# syntax = docker/dockerfile:experimental

FROM --platform=${BUILDPLATFORM} golang:1.14 AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

FROM golang:1.14 AS build
WORKDIR /opt
ARG TARGETOS
ARG TARGETARCH
COPY . .
RUN mkdir /out && CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/01cloud-payments cmd/server/main.go

FROM base AS unit-test
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    mkdir /out && go test -v -coverprofile=/out/cover.out ./...

FROM golangci/golangci-lint:v1.31.0-alpine AS lint-base

FROM base AS lint
RUN --mount=target=. \
    --mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci-lint \
    golangci-lint run --timeout 10m0s ./...


FROM scratch AS unit-test-coverage
COPY --from=unit-test /out/cover.out /cover.out

FROM scratch AS bin-unix
COPY --from=build /out/01cloud-payments /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/01cloud-payments /01cloud-payments.exe

FROM bin-${TARGETOS} as bin

FROM alpine:latest as final
WORKDIR /app 
RUN touch .env
COPY --from=build /out/01cloud-payments .
CMD ["./01cloud-payments"]
