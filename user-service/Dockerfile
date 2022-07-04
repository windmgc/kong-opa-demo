FROM golang:1.18-alpine as builder
RUN mkdir /go/src/user-service
RUN apk add git
COPY . /go/src/user-service
WORKDIR /go/src/user-service
RUN go get -v ./...
RUN go build -o /go/bin/user-service .

FROM alpine:latest
RUN mkdir /user-service
WORKDIR /user-service
COPY --from=builder /go/bin/user-service /user-service/user-service
RUN chmod 755 /user-service/user-service
ENTRYPOINT ["/user-service/user-service"]
