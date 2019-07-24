FROM golang:alpine as builder

WORKDIR /build/action
ADD . /build/action

# https://github.com/google/go-github/issues/1049
RUN apk update > /dev/null 2>&1 && apk upgrade > /dev/null 2>&1 && apk add --no-cache gcc musl-dev git ca-certificates > /dev/null 2>&1

RUN CGO_ENABLED=0 GOOS=linux \
    go get -v ./... > /dev/null 2>&1 && \
    go test -v ./... > /dev/null 2>&1 && \
    go build -o build/email ./cmd/ > /dev/null 2>&1

FROM alpine:latest

RUN apk update && \
      apk upgrade && \
      apk add --no-cache ca-certificates && \
      update-ca-certificates

LABEL version="1.0.0"
LABEL repository="https://github.com/cirrus-actions/actions-trigger/"
LABEL homepage="https://github.com/marketplace/cirrus-ci"
LABEL maintainer="Cirrus Labs"
LABEL "com.github.actions.name"="Email"
LABEL "com.github.actions.description"="Emails check suite results upon completion"
LABEL "com.github.actions.icon"="mail"
LABEL "com.github.actions.color"="green"

COPY --from=builder /build/action/build/email /actions/email
ENTRYPOINT exec /actions/email
