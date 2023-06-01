FROM golang:1.18-alpine3.15 AS builder
RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:3.18.0

WORKDIR /

COPY --from=builder /src/main /main
EXPOSE 8081
CMD /main