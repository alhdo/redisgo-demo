FROM golang:1.21

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD [ "go", "run", "/app/cmd/simple-worker/main.go", "-c", "/app/simple-go-worker.conf.example" ]