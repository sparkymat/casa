FROM hub.orion.home/golang:1.19-alpine AS builder

COPY . /app/

WORKDIR /app
RUN CGO_ENABLED=0 go build -o casa .

FROM hub.orion.home/alpine:3

COPY --from=builder /app/casa /bin/casa

CMD ["/bin/casa"]
