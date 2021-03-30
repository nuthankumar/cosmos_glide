FROM alpine:3.10
RUN apk add --no-cache \
        libc6-compat
COPY ./bin/cosmos_crd /app
ENV CQL_TIMEOUT_MS=60000
RUN chmod +x app
ENTRYPOINT /app
