FROM alpine:latest

ARG BUILD_PATH

COPY ${BUILD_PATH} /server

CMD ["/server"]
