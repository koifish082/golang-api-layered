FROM golang:1.15.0-alpine

# Required to fetch private Github repos
ARG SSH_KEY

RUN apk update \
    && apk add openssh-client git \
    && mkdir -p /root/.ssh && chmod 0700 /root/.ssh \
    && ssh-keyscan github.com >/root/.ssh/known_hosts

# Move to project root
WORKDIR /go/src/github.com/koifish082/golang-api-layered
COPY main.go go.mod ./
COPY cmd cmd
COPY src src
COPY mock mock

RUN echo "$SSH_KEY" > /root/.ssh/id_rsa \
    && chmod 0600 /root/.ssh/id_rsa \
    && go install . \
    && rm /root/.ssh/id_rsa

CMD [ "golang-api-layered", "serve" ]
