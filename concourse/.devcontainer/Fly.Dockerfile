FROM alpine:latest as resource

ENV CONCOURSE_VERSION=6.7.0

RUN apk add --update --no-cache \
    ca-certificates \
    ncurses \
    make \
    git \
    curl \
    tar
RUN curl -sL https://github.com/concourse/concourse/releases/download/v${CONCOURSE_VERSION}/fly-${CONCOURSE_VERSION}-linux-amd64.tgz \
    | tar -xz  -C /usr/local/bin && \
    curl -sL https://taskfile.dev/install.sh | sh