FROM ubuntu:18.04

ARG SHORT_SHA

RUN apt-get update

COPY our-binary-$SHORT_SHA /our-binary

RUN chmod 755 /our-binary

ENTRYPOINT /our-binary

