# syntax=docker/dockerfile:1.2
FROM golang:1.18-alpine3.16 AS build

LABEL MAINTAINER="kh411d"

RUN apk add --no-cache ca-certificates git curl wget zip build-base openssh-client

WORKDIR /build

COPY . .
RUN --mount=type=ssh \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build,id=goshrpac \
    go mod download \
  && CGO_ENABLED=0 go build -v -installsuffix 'static' -o ./app ./app

######

FROM alpine:3.16

RUN apk add --no-cache ca-certificates curl zip wget tzdata

# files from GOROOT
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
# build result
COPY --from=build /build/app/app /app

ENTRYPOINT ["/app","httpd"]

EXPOSE 3000
