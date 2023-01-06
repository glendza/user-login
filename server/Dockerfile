FROM golang:1.19

WORKDIR $GOPATH/src/glendza/login-code-challenge
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 5678

CMD ["login-code-challenge"]