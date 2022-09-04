FROM golang:latest

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN go install github.com/golang/mock/mockgen@v1.5.0

WORKDIR /opt/app

CMD ["air"]