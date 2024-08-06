FROM golang:1.22-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o authenticationService .

RUN chmod +x /app/authenticationService

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/authenticationService /app/id_rsa_priv.pem /app/

CMD [ "/app/authenticationService" ]