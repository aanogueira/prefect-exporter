FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum *.go /app/

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o prefect-exporter .

RUN echo 'nobody:x:65534:65534:nobody:/:' > /app/passwd && \
    echo 'nobody:x:65534:' > /app/group

RUN apk --no-cache add ca-certificates

FROM scratch

WORKDIR /app

COPY --from=builder /app/prefect-exporter /app/
COPY --from=builder /app/group /app/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER nobody:nobody

ENTRYPOINT [ "/app/prefect-exporter" ]
