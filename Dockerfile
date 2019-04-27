FROM golang:alpine

WORKDIR /app
COPY main.go /app
CMD ["go","run","main.go"]