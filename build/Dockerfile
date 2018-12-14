FROM alpine:3.8

RUN apk upgrade --update --no-cache

USER nobody

ADD build/_output/bin/VTN-operator /usr/local/bin/vtn-operator
