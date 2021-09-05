

FROM golang:latest

RUN mkdir -p /go/src/github.com/user/app/
COPY . /go/src/github.com/user/app/
WORKDIR /go/src/github.com/user/app/
COPY . .
COPY go.mod ./
COPY go.sum ./
RUN go mod download




CMD ["go", "run", "main.go"]