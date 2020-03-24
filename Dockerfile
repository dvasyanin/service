FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/mindbox-srv-go

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/*.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/mindbox-srv-go/main .
COPY --from=builder /app/mindbox-srv-go/config.yaml .

CMD ["./main"]
