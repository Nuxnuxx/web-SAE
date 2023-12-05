FROM golang:1.21-alpine

RUN apk add --no-cache make

RUN adduser -D backend
RUN echo 'backend:password' | chpasswd
USER backend

WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go .
COPY Makefile .
RUN make build

CMD ["/app/bin/backend"]
