FROM golang:1.17-alpine
ADD . /go/src/hello-telemetry
RUN go install hello-telemetry

FROM alpine:latest
COPY --from=0 /go/bin/hello-telemetry .
ENV PORT 8080
CMD ["./hello-telemetry"]