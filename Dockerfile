FROM golang:1.22.4-alpine
WORKDIR /src
COPY . .
RUN go build -o /serverbin ./server/server.go
ENTRYPOINT [ "/serverbin" ]