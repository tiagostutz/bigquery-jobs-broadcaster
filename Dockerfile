# Anything that we need in to build (e.g. swagger, protoc, git) goes here.
FROM golang:1.18-alpine AS build_base
RUN apk add --no-cache protobuf-dev git

FROM build_base AS build

WORKDIR /source/
COPY . .

RUN CGO_ENABLED=0 go install -v
RUN CGO_ENABLED=0 go build -v -o /go/bin/app

FROM gcr.io/distroless/static

COPY --from=build /go/bin/app /app
WORKDIR /

ENTRYPOINT ["/app"]

