# Dockerfile is using in production.

# Base stage for getting go-related dependencies.
FROM golang:alpine as base_stage
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/eeonevision/avito-pro-test
ENV GO111MODULE=on

# Go module is a magic :)
COPY go.mod .
COPY go.sum .
RUN go mod download

# Compile stage.
FROM base_stage as compile_stage
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/avitornd

# Fresh alpine uses to reduce image size and not ship sources, Go compiler.
FROM alpine AS avitornd
COPY --from=compile_stage /go/bin/avitornd /bin/avitornd
ENTRYPOINT ["/bin/avitornd"]
