FROM golang:1.12

RUN mkdir -p /ozone/platform
ADD . /ozone/platform
WORKDIR /ozone/platform

RUN go get -d
RUN go build -o main .
CMD ["/ozone/platform/main"]