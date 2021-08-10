FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod /app
COPY main.go /app

RUN go mod download
RUN go build -o /myapp

CMD /myapp

EXPOSE 8000