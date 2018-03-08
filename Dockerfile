FROM golang:1.9.4-alpine3.7 AS build

WORKDIR /go/src/github.com/avegao/iot-openevse-service

RUN apk add --no-cache --update \
    git \
    glide

COPY ./glide.yaml ./glide.lock ./

RUN glide install

COPY ./ ./

ARG VCS_REF="unknown"
ARG BUILD_DATE="unknown"

RUN go install

########################################################################################################################

FROM alpine:3.7

MAINTAINER "Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

WORKDIR /app

RUN addgroup iot-openevse-service && \
    adduser -D -G iot-openevse-service iot-openevse-service && \
    chown iot-openevse-service:iot-openevse-service -R ./

USER iot-openevse-service

COPY --from=build /go/bin/iot-openevse-service /app/iot-openevse-service
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

EXPOSE 50000/tcp

LABEL com.avegao.iot.openevse.vcs_ref=$VCS_REF \
      com.avegao.iot.openevse.build_date=$BUILD_DATE \
      maintainer="Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

VOLUME /app/export

ENTRYPOINT ["./iot-openevse-service"]
