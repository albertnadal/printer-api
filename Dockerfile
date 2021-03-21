FROM golang:1.15-alpine
MAINTAINER Albert Nadal Garriga
RUN mkdir /printer-api
ADD . /printer-api
WORKDIR /printer-api
RUN apk add --update alpine-sdk
RUN go build -o server
EXPOSE 8080
CMD /printer-api/server
