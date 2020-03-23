FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o merchant-client

# RUN

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /APP
ADD merchant.json /app/merchant.json
COPY --from=builder /app/merchant-client .

CMD ["./merchant-client"]