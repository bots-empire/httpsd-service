FROM golang:1.18

WORKDIR /app

COPY . .
RUN go build -buildvcs=false -o /bin/httpsd-service ./cmd/httpsd

CMD ["/bin/httpsd-service"]
