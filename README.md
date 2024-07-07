# DateTime-Server-Rowan
A basic HTTP server that returns the current date and time.

## 1. Server Functionality:

- Endpoint: `GET /datetime`
- Response: Current date and time

## 2. Implementations:

The server is implemented using:
1. Standard library `(net/http)`
2. `Gin Framework`

## 3. Directories Walkthrough
        cmd  
        |____httpserver
        |              |___server.go (main code)
        |              |___server_test.go (test file)
        |
        |____ginserver
                      |___server.go (main code)
                      |___server_test.go (test file)
## 4. Ports 
   - HTTP server on `port 8080`
   - HTTP using Gin on `port 8089`

## 5. Docker Guide
### Docker Files:
- Both files are using multistage build to enhance and optimize space
   - `Dockerfile-http`
       - 311MB    -----> before multistage build
       - 7MB    -----> after multistage build
   - `Dockerfile-gin`
       - 569MB    -----> before multistage build
       - 10.8MB    -----> after multistage build

### Docker Compose
- Includes 2 services
- Both services can be accessed either:
  - Explicit unique name of image: 
    - `datetime-server-rowan-httpserver`
    - `datetime-server-rowan-ginserver`
  - Using build of dockerfiles
- Ports are mapped accordingly as stated before (8080:8080) & (8089:8089)

### Docker commands: (use sudo if needed)

```
docker build --tag=datetime-server-rowan-httpserver . -f Dockerfile-http
```
```
docker build --tag=datetime-server-rowan-ginserver . -f Dockerfile-gin
```
```
docker run -p 8080:8080 datetime-server-rowan-httpserver
```
```
docker run -p 8089:8089 datetime-server-rowan-ginserver
```
```
docker compose up
```
## 6. Makefile
### Targets:
1. build
2. test
3. run
4. clean

### Commands: 

```
make //automatically builds
```
```
make test
```
```
make run
```
```
make clean
```

### To Send Requests:
```
curl localhost:8080/datetime
```
```
curl localhost:8089/datetime
```
#### NOTE: For testing without docker compose or specifying port when using image, substitute localhost by IP of container.