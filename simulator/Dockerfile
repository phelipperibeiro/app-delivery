FROM golang:1.16

## WORKDIR -> folder that we will work inside the container
WORKDIR /go/src

## environment variable
ENV PATH="/go/bin:${PATH}"

## librdkafka-dev -> library that allows us to work with kafka
RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

## CMD -> I used this just to keep the process active, that is, keep the container running
CMD ["tail", "-f", "/dev/null"]