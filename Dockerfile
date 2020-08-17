FROM golang:1.14-alpine as builder

WORKDIR /go/src/app

COPY . .

RUN go get -d -v
# RUN go build -o main .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -ldflags='-w -s -extldflags "-static"' -a \
  -o main .

FROM scratch

COPY --from=builder /go/src/app/main /go/main

ENTRYPOINT [ "/go/main" ]
