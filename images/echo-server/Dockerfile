FROM golang:1.21
# https://gist.github.com/olivernadj/e1a02f5ed1731c2ccced554cec0fc928
ENV CGO_ENABLED=0
COPY main.go ./
# This file is to force a different image for testing.
RUN echo "$(date)" > /build_time.txt
RUN go build -o /echo-server main.go
ENTRYPOINT [ "/echo-server" ]