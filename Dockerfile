FROM golang:1.17-alpine as builder

LABEL maintainer="Alexander Miranda <alexandermichaelmiranda@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN apk --no-cache add tzdata

RUN go mod download

COPY . .

RUN go build -o main /app/cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app .

EXPOSE 8080:80

CMD ["./main"]
