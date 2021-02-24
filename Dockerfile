FROM golang:alpine

WORKDIR /go/src/app

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "api/main.go"]