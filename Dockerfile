FROM golang:1.22rc1-alpine3.19 as go-builder

USER root

ARG GOPRIVATE
ARG GITHUB_TOKEN
ARG GITHUB_USERNAME

RUN go env -w CGO_ENABLED=1
RUN go env -w GOPRIVATE=$GOPRIVATE
RUN go env -w GOCACHE="/root/.cache/go-build"
RUN go env -w GOMODCACHE="/root/.cache/go-build"


RUN apk add --no-cache git ca-certificates

RUN git config --global url."https://$GITHUB_USERNAME:$GITHUB_TOKEN@github.com/".insteadOf "https://github.com/"

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target="/root/.cache/go-build/" go mod download -x

COPY . .
RUN --mount=type=cache,target="/root/.cache/go-build/" go build -tags musl ./cmd/main-service

FROM alpine:latest

WORKDIR /app

COPY --from=go-builder /app/main-service /usr/local/bin

ENTRYPOINT main-service --env=$ENVIRONMENT --host=$HOST --port=$PORT
