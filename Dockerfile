FROM golang:1.21

WORKDIR /usr/src/facutured-reception-backend

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/Go_Gin ./...

CMD ["Go_Gin"]
