FROM golang:1.22

WORKDIR /api

COPY . .

RUN go mod download & go mod verify

RUN CGO_ENABLED=1 GOOS=linux go build -v -o api

EXPOSE 8080

CMD ["./api"]