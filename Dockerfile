# builder image
FROM golang:1.18.1-alpine as base
LABEL stage=builder
WORKDIR /builder
RUN apk add upx
RUN apk add make
ENV GO111MODULE=on CGO_ENABLED=0
COPY . .
RUN make tidy
RUN make build

# runner image
FROM alpine:3.8
WORKDIR /app
COPY --from=base /builder/go-test-service go-test-service
COPY --from=base /builder/config.yml config.yml
CMD ["/app/go-test-service", "server"]