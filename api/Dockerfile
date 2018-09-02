FROM golang:1.11-alpine3.7

ENV APPDIR /go/src/github.com/Khigashiguchi/khigashiguchi.com/api
ENV DEP_VERSION 0.5.0

WORKDIR ${APPDIR}
COPY . ${APPDIR}

RUN apk add --no-cache ca-certificates \
    dpkg \
    gcc \
    git \
    musl-dev \
    openssh \
    bash \
    curl

RUN curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o /go/bin/dep
RUN chmod +x /go/bin/dep
RUN /go/bin/dep ensure

RUN go build -v -o api

EXPOSE 8080
CMD ["./api"]