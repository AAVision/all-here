# all-here

<p align="center">
  <img width="50%" height="300" src="all-here.png">
</p>

We always face some obstacles while testing the vendor's endpoints from the backend side and even from the front-end team. So we decided to create a tool that will return to you the status code you want. The tool in default runs on port 9876, but you can pass the port you want.

## HTTP status codes

HTTP status codes are three-digit responses from the server to the browser-side request. Everyone has probably gotten the classic 404 page-not-found error. That is an HTTP client error status code and there are a lot more of them.
https://umbraco.com/knowledge-base/http-status-codes/
https://github.com/golang/go/blob/master/src/net/http/status.go

## Info
You can get the release: https://github.com/AAVision/all-here/releases/tag/v0.0.1

## Run
```bash
go run . --port 9887
```
then try to hit the following endpoint with the status code you want:
```bash
http://localhost:9887/504
```

## Build 
```bash
go build .
```

## LICENSE

This project is licensed under the MIT License. See the [LICENSE](https://github.com/aavision/all-here/blob/main/LICENSE) file for details.
