FROM hub.orion.home/golang:1.19-alpine AS builder

COPY . /app/

WORKDIR /app
RUN go generate ./...
RUN make casa

FROM hub.orion.home/alpine:3

COPY --from=builder /app/casa /bin/casa

WORKDIR /app
COPY public /app/public

CMD ["/bin/casa"]
