FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -tags dev -o main api/main.go

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=builder /dist/main /

EXPOSE 8080

ENTRYPOINT ["/main"]
