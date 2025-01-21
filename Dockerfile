# Build
FROM golang:1.23-alpine3.21 AS build

RUN mkdir /go/src/app
COPY . /go/src/app
WORKDIR /go/src/app

# RUN apk update
# RUN apk add -U --no-cache ca-certificates && update-ca-certificates
# RUN apk --no-cache add ca-certificates

# RUN go mod tidy
RUN CGO_ENABLED=0 go build -mod vendor -trimpath -ldflags="-s -w" -o /go/src/app/main main.go

# Run
# FROM chromedp/headless-shell:latest
FROM scratch

# COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/app/main /go/src/app/

WORKDIR /go/src/app
EXPOSE 8090

# RUN apt-get update; apt install dumb-init -y
# ENTRYPOINT ["dumb-init", "--"]

CMD ["/go/src/app/main"]


