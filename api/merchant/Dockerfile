FROM golang:alpine as builder

ARG goprivate
ARG gitlab_user
ARG gitlab_token


RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV gitlab_user=$gitlab_user
ENV gitlab_token=$gitlab_token
ENV GOPRIVATE=$goprivate
ENV GO111MODULE=on
ENV GIT_TERMINAL_PROMPT=1
ENV CGO_ENABLED=0
ENV GOOS=linux

COPY . .

RUN git config \
    --global \
    url."https://${gitlab_user}:${gitlab_token}@gitlab.com/".insteadOf \
    "https://gitlab.com/"

RUN env

RUN go build -a -installsuffix cgo -o api-merchant

# RUN

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /APP
COPY --from=builder /app/api-merchant .

CMD ["./api-merchant"]