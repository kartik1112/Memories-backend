FROM golang:1.22-alpine

RUN apk update && apk add bash
RUN apk add --no-cache tzdata


WORKDIR /app

COPY go.* ./

COPY wait-for-it.sh /wait-for-it.sh

RUN go mod tidy

COPY . .

RUN chmod +x /wait-for-it.sh

RUN go build -o main main.go

EXPOSE 3000

CMD ["./main"]