FROM golang:1.22-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o loggerService .

RUN chmod +x /app/loggerService

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/loggerService /app

CMD [ "/app/loggerService" ]