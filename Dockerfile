FROM golang:1.7.1-alpine

ENV BUILD_PACKAGES git

# Update and install all of the required packages.
# At the end, remove the apk cache
RUN apk update && \
    apk upgrade && \
    apk add $BUILD_PACKAGES && \
    rm -rf /var/cache/apk/*

# Copy the source
ADD . /go/src/github.com/mloudon/mcdemo

WORKDIR /go/src/github.com/mloudon/mcdemo

RUN go install .

EXPOSE 5000

CMD ["/go/bin/mcdemo"]
