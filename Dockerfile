FROM golang:1.23-bookworm AS builder

ARG UPX_VERSION=4.2.4

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends apt-utils xz-utils && \
    curl -Ls https://github.com/upx/upx/releases/download/v${UPX_VERSION}/upx-${UPX_VERSION}-amd64_linux.tar.xz -o - | tar xvJf - -C /tmp && \
    cp /tmp/upx-${UPX_VERSION}-amd64_linux/upx /usr/local/bin/ && \
    chmod +x /usr/local/bin/upx && \
    apt-get remove -y xz-utils && \
    rm -rf /var/lib/apt/lists/* && \
    useradd -u 10001 app && \
    touch config.env

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make build && \
    make compress && \
    make test-binary


FROM scratch AS final

ARG WEB_PORT=8080

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /app/config.env .

COPY --from=builder /app/build .

EXPOSE ${WEB_PORT}

USER app

ENTRYPOINT [ "./app" ]
