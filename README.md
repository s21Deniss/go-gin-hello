# This is simple application to test consul

This application is made to lab hashicorp consul

Can change app port by setting env variable PORT to desired port `export PORT=8181`

## Routes

* 	"/" : `prints "Hello World!"`

* "/health" : `returns 'status' 200 or 503`

* "/setOK" : `sets status to ok and return 200 from /health`

* "/setFail" : `sets status to fail and return 503 from /health`


# Running localy

fetch modules : `go get -d -v`
run : `go run main.go`

# Docker image

Building your self : `docker build -t 'selfbuild/go-gin-hello' .`

OR

Pull from dockerhub : `docker pull s21deniss/go-gin-hello`


---

Start it

Run it : `docker run -p 8181:8181 -e PORT=8181 s21deniss/go-gin-hello`

---
To register in consul you should have consul cli installed/configured
Register in consul ` make consul grepSearch=s21deniss/go-gin-hello `
