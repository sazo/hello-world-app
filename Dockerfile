FROM golang:1.17 as builder

ADD . .
RUN go build main.go

FROM gcr.io/distroless/base

COPY --from=builder /go/main /
CMD [ "/main" ]