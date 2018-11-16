# build stage
FROM golang:1.11.2-alpine3.8 AS builder

RUN apk update && \
    apk add libressl-dev bash git ca-certificates curl

# establish a working directory
WORKDIR /go/src/github.com/alecholmez/http-server/
ADD . /go/src/github.com/alecholmez/http-server/

# vendor dependencies in the builder container
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    cd /go/src/github.com/alecholmez/http-server && \
    dep ensure --vendor-only -v

# build our binary
RUN cd h2tp && \
    go build .

# final stage
FROM alpine:3.8

RUN apk update && \
    apk add ca-certificates

WORKDIR /app

# Copy over the neccessary items from the builder container: binary, entrypoint script, and certs
COPY --from=builder /go/src/github.com/alecholmez/http-server/h2tp/h2tp /app/

CMD ./h2tp

EXPOSE 6060
EXPOSE 9000
