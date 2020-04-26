FROM golang:alpine as builder

ARG repo_name
ARG gitlab_user
ENV gitlab_user=$gitlab_user
ARG gitlab_token
ENV gitlab_token=$gitlab_token

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on
ENV GOPRIVATE=${repo_name}
ENV GIT_TERMINAL_PROMPT=1
ENV CGO_ENABLED=0
ENV GOOS=linux

COPY . .

RUN git config \
    --global \
    url."https://${gitlab_user}:${gitlab_token}@gitlab.com/".insteadOf \
    "https://gitlab.com/"

RUN go build -a -installsuffix cgo -o api-merchant

# RUN

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /APP
COPY --from=builder /app/api-merchant .

CMD ["./api-merchant"]