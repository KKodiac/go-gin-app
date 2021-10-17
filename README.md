# go-gin-app

Simple web app using Go and Gin framework

Golang 과 Gin 프레임워크를 사용한 간단한 웹 앱

## How to get Started
Install Gin and have Go installed on your system.
- Installing Gin
```sh
    go get -u github.com/gin-gonic/gin
```
- Building executable
```sh
    go build -o app
```
- Running the executable
```sh
    ./app
    # Check on localhost:8080
```
- Testing
```sh 
    go test -v
```

## Fetching and Testing
Responses are set according to `Accept` headers
- HTML
```sh
    # Retrievs List of Articles in HTML format
    curl -X GET -H "Accept: application/html" http://localhost:8080/
``` 
- JSON
```sh
    # Retrievs Content of Articles 1 in JSON format
    curl -X GET -H "Accept: application/json" http://localhost:8080/article/view/1
```
- XML
```sh
    # Retrievs Content of Articles 2 in JSON format
    curl -X GET -H "Accept: application/xml" http://localhost:8080/article/view/2
```