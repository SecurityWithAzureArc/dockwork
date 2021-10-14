FROM golang:1.16-alpine AS build-init

# Install UPX
ARG UPX_TAR_NAME=upx-3.95-amd64_linux
RUN apk --no-cache add git && \
    wget -O ${UPX_TAR_NAME}.tar.xz https://github.com/upx/upx/releases/download/v3.95/$UPX_TAR_NAME.tar.xz && \
    tar -xf $UPX_TAR_NAME.tar.xz && \
    rm -f $UPX_TAR_NAME.tar.xz && \
    mv $UPX_TAR_NAME/upx /usr/local/bin && \
    rm -rf $UPX_TAR_NAME

FROM build-init as builder

WORKDIR /app

# Install Deps
COPY go.mod go.sum ./
RUN go mod download

# Build App
COPY . .
RUN CGO_ENABLED=0 go generate ./...
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /out/app .

# Compress App
RUN upx /out/*

FROM scratch

# Since we started from scratch, we'll copy the SSL root certificates from the builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /out/* /usr/local/bin/

ENTRYPOINT [ "app" ]
