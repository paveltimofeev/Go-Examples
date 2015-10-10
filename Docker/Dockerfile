# https://hub.docker.com/_/golang/
# https://github.com/docker-library/golang/blob/51d6eacd41fe80d41105142b9ad32f575082970f/1.5/Dockerfile

FROM golang:1.5.1

ENV CHANGED 2015-10-10:v1

RUN \
cd /go/src/ && \
mkdir github.com && cd github.com &&\
mkdir paveltimofeev && cd paveltimofeev && \
git clone https://github.com/paveltimofeev/Go-Examples.git

WORKDIR /go/src/github.com/paveltimofeev/Go-Examples

CMD godoc -http=":8080"

EXPOSE 8080 
