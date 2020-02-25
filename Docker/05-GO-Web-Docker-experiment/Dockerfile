#Go Docker file
FROM golang:1.13
WORKDIR $GOPATH/src/github.com/rpfeb23/rakesh-pattanayak-webdev-go/Docker/05-GO-Web-Docker-experiment

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["05-GO-Web-Docker-experiment"]
MAINTAINER rpfeb23@gmail.com
