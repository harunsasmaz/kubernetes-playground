FROM golang:1.17.3-alpine

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server cmd/todo/main.go

FROM alpine
ENV PORT=8000

COPY --from=0 /src/server /usr/bin/server
CMD ["server"]