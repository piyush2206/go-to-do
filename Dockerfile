FROM golang:1.10.2-alpine
COPY . /go/src/github.com/piyush2206/go-to-do
WORKDIR /go/src/github.com/piyush2206/go-to-do
RUN go build -o /go/bin/app .

FROM alpine:3.5
LABEL Author=piyush.vishwakarma
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/app .
CMD ["/app"]
EXPOSE 80