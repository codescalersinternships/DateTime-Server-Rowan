build: 
	go build -o httpserver ./cmd/httpserver/server.go
	go build -o ginserver ./cmd/ginserver/server.go

test: build
	go test ./...

run compose: build
	sudo docker compose up

clean:
	rm httpserver
	rm ginserver
fmt:
	go fmt ./...