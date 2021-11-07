FROM golang:alpine
RUN apk add --no-cache git
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -tags prod -o /app/magiclink /app/api/server.go
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["/app/magiclink"]
