FROM golang:1.16-alpine

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./backend

EXPOSE 8090

CMD ["./backend"]