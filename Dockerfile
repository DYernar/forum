FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o main .
CMD ["/app/main"]
