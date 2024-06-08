FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o authenticationService main.go

RUN chmod +x /app/authenticationService

CMD [ "/app/authenticationService" ]

# TODO: fix issue with with copying private key for small docker image