FROM golang as builder

WORKDIR /go/src/kutala

ENV GOOS=linux
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOARCH=amd64

RUN go mod init github.com/Paulo-Lopes-Estevao/kutala-api-gateway

COPY . /go/src/kutala/

RUN go build -o main

EXPOSE 9000

RUN chmod +x ./main

CMD ["./main"]