FROM golang:1.11.1

WORKDIR /go/src/apodExample
COPY . .

EXPOSE 8080

RUN go get -d -v ./...
RUN go install -v ./...

ARG nasa_api_key=''

ENV NASA_API=$nasa_api_key

ENTRYPOINT [ "apodExample" ]