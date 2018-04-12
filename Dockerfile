# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.8

COPY ./source /go/src/to-do-list/source

WORKDIR /go/src/to-do-list/source

RUN go get -v && \
go build main.go

RUN pwd && \
ls -lah

CMD [ "./main" ]

