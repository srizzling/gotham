FROM golang:1.8 as builder
ARG SERVICE
LABEL maintainer Sriram Venkatesh <venksriram@gmail.com>
COPY ./ /go/src/github.com/srizzling/gotham
WORKDIR /go/src/github.com/srizzling/gotham/$SERVICE
RUN go-wrapper download
RUN go-wrapper install
RUN go build -a -installsuffix cgo -ldflags '-w' -o $SERVICE

# Run binary container
FROM scratch
ARG SERVICE
WORKDIR /root/
COPY --from=builder /go/src/github.com/srizzling/gotham/$SERVICE .
ENTRYPOINT [$SERVICE]
