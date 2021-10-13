FROM golang:1.16-alpine

WORKDIR /app

# Install tools
RUN apk add --no-cache git && \
    GO11MODULE=off go get github.com/cosmtrek/air

# Install deps
COPY go.* /app/
RUN go mod download

ADD . .

EXPOSE 5000
CMD ["air"]
