FROM golang:1.15 AS base

WORKDIR /go/src/app
COPY . .

RUN go build main.go

COPY --from=base main /bin/main

CMD /bin/main