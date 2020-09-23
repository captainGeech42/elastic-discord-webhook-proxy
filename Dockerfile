# IMAGE 1: Builder
FROM golang:1.14.4-alpine as builder

# pre-reqs
RUN apk add --no-cache git

# copy in src
WORKDIR $GOPATH/src/elastic-discord-webhook-proxy/
COPY *.go ./

# install dependencies
RUN go get -d -v

# build
RUN go build -o /proxy

# IMAGE 2: Runner
FROM alpine:latest

# copy binary
WORKDIR /app
COPY --from=builder /proxy .

# run
EXPOSE 8080
CMD /app/proxy
