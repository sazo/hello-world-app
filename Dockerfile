FROM golang:1.17 as builder

ADD . .
RUN go build main.go

FROM centos:7

COPY --from=builder /go/main /
CMD [ "/main" ]