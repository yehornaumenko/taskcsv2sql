FROM golang:1.15 AS base

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 go build main.go

FROM alpine:3.5
COPY --from=base /go/src/app/main /bin/main
COPY --from=base /go/src/app/config.yaml /config_file.yaml

RUN chmod +x /bin/main
CMD /bin/main --config=/config_file.yaml