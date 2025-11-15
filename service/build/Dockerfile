FROM golang:1.23.0

ENV APPPATH=$GOPATH/src/service
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

WORKDIR $APPPATH

COPY ./ .

# Update Debian packages to reduce vulnerabilities
RUN apt-get update && apt-get upgrade -y

RUN go mod tidy && go mod download

RUN go build -o /main ./service/cmd/main.go

EXPOSE 8080

CMD ["/main"]
