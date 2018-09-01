FROM golang:1.11-alpine3.7

ENV APPDIR /go/src/github.com/Khigashiguchi/khigashiguchi.com/api
WORKDIR ${APPDIR}
COPY . ${APPDIR}

RUN ls -la

RUN apk add --no-cache ca-certificates \
    dpkg \
    gcc \
    git \
    musl-dev \
    openssh \
    bash \
    curl

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

RUN go build -v -o api

EXPOSE 8080
CMD ["./api"]