FROM golang:1.22-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o gatewayService .

RUN chmod +x /app/gatewayService

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/gatewayService /app/private-key.pem /app/

CMD [ "/app/gatewayService" ]