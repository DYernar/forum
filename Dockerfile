FROM golang:latest
WORKDIR /usr/src/app/
COPY . /usr/src/app/
RUN go mod download
RUN go build -o forum .
EXPOSE 8001
ENV TZ Asia/Almaty
CMD ["./forum"]