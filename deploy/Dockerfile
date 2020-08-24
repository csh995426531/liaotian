FROM alpine:latest AS user-service
WORKDIR /
COPY ./user-service/main ./
CMD [ "./main" ]

FROM alpine:latest AS user-web
WORKDIR /
COPY ./user-web/main ./
CMD [ "./main" ]