FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk update
RUN apk upgrade
RUN apk add --update go=1.13.10-r0 gcc=9.2.0-r4 g++=9.2.0-r4
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o forum_unix -v
CMD ["/app/forum_unix"]

