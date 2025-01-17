# Build
FROM golang:latest AS build

RUN mkdir /go/src/app
COPY . /go/src/app

WORKDIR /go/src/app

RUN go mod tidy
RUN go build -o /go/src/app/main main.go

# Run
FROM chromedp/headless-shell:latest

COPY --from=build /go/src/app/main /go/src/app/
# COPY --from=build /go/src/app/templates /go/src/app/templates

WORKDIR /go/src/app

EXPOSE 8090

RUN apt-get update; apt install dumb-init -y

ENTRYPOINT ["dumb-init", "--"]
CMD ["/go/src/app/main"]