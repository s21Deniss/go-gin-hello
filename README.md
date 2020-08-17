# This is simple application to test consul

This application is made to lab hashicorp consul

## Routes

* 	"/" : `prints "Hello World!"`

* "/health" : `returns 'status' 200 or 503`

* "/setOK" : `sets status to ok and return 200 from /health`

* "/setFail" : `sets status to fail and return 503 from /health`


# Docker image

Building your self : `docker build -t 'selfbuild/go-gin-hello'`

Pull from dockerhub : `docker pull s21deniss/go-gin-hello`
