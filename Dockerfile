FROM golang:1.18-buster

RUN mkdir /app

ADD . /app

WORKDIR /app


RUN go mod tidy

RUN go mod download

RUN go clean --modcache

COPY go.mod .

COPY go.sum .

COPY . .

RUN go build -o app.out

CMD ["./run.sh"]
