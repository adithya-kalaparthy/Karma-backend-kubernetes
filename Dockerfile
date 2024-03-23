FROM golang:1.22.0-alpine3.18

WORKDIR /karma

COPY .env .

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "cmd/server/main.go"]

