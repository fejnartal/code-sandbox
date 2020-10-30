FROM buildpack-deps:buster-scm

RUN apt-get update && apt-get install -y \
  git \
  icu-devtools \
  && rm -rf /var/lib/apt/lists/*
