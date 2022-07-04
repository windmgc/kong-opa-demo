FROM golang:1.18-alpine as builder
RUN mkdir /go/src/article-service
RUN apk add git
COPY . /go/src/article-service
WORKDIR /go/src/article-service
RUN go get -v ./...
RUN go build -o /go/bin/article-service .

FROM alpine:latest
RUN mkdir /article-service
WORKDIR /article-service
COPY --from=builder /go/bin/article-service /article-service/article-service
RUN chmod 755 /article-service/article-service
ENTRYPOINT ["/article-service/article-service"]
