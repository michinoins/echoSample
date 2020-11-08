###############################
# Builder Container
###############################
FROM golang:1.15.0 AS builder

# Go Moduleが使えるようにする
ENV GO111MODULE=on

WORKDIR /go/src/echoSample/

# GOROOTが/goになるため
COPY ./src /go/src/echoSample
RUN go get github.com/pilu/fresh \
    && go get github.com/labstack/echo \
    && go get github.com/labstack/echo/middleware

EXPOSE 8080
CMD ["fresh"]

###############################
# Production Container
###############################
FROM golang:1.15.0-alpine AS production
COPY --from=builder /go/1.15.0/src/echoSample/ /echoSample