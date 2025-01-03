FROM golang:1.23

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /test_webserver

CMD [ "/test_webserver" ]