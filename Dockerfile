FROM golang:1.20.6

WORKDIR /BlogifyAPI
COPY go.mod go.sum ./
COPY go.* ./ 
RUN go mod download
COPY *.go ./
RUN go build -o blogifyapi

EXPOSE 8080

CMD ["./blogifyapi"]