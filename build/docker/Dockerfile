# STEP 1 build executable binary

FROM golang:1.11.1-alpine3.8 as builder

# Install SSL ca certificates & create user
RUN apk add --update --no-cache git ca-certificates \
    && adduser -D -g '' appuser

COPY src $GOPATH/src/smartinsurtech.de/echo-adapter/
WORKDIR $GOPATH/src/smartinsurtech.de/echo-adapter/

#get app dependencies and build the binary
RUN go get -d -v \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o /go/bin/main

# STEP 2 build a small runtime image

# start from scratch
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static executable
COPY --from=builder /go/bin/main /go/bin/main

USER appuser

EXPOSE 8080