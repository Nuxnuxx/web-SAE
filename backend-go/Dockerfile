FROM golang:1.21-alpine
RUN apk add --no-cache make

RUN adduser -D backend
RUN echo 'backend:password' | chpasswd
USER backend

WORKDIR /app/

COPY ./api/ ./api/
COPY ./config/ ./config/
COPY ./storage/ ./storage/
COPY ./types/ ./types/
COPY ./utils/ ./utils/
COPY main.go ./
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY Makefile .
RUN make build

CMD ["/app/bin/backend"]
