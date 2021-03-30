FROM alpine:3.10
RUN apk add --no-cache \
        libc6-compat
COPY ./bin/ccspaas_crd /app
ENV NAMESPACE=default
ENV SERVICE=default
ENV SECRET_NAME=default
ENV CQL_TIMEOUT_MS=60000
RUN chmod +x app
ENTRYPOINT /app
