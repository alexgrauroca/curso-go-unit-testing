ARG GO_VERSION=1.18.1

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /src

COPY main.go main.go
COPY ./go.mod ./go.sum ./

RUN go mod vendor

COPY util util
COPY controller controller
COPY models models

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /gounittesting

FROM scratch AS runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /gounittesting /go-unit-testing

EXPOSE 8080

ENTRYPOINT [ "/go-unit-testing" ]
