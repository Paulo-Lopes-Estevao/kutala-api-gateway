FROM golang as builder

WORKDIR /go/src/kutala

ENV GOOS=linux
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOARCH=amd64

COPY . /go/src/kutala/


RUN go build -o main.go

WORKDIR /go/src/kutala

FROM scratch

COPY --from=builder /go/src/kutala/ /
COPY --from=builder /go/src/kutala/ ./
#COPY --from=builder /go/src/kutala/main ./main
COPY --from=builder /go/src/kutala/app /app
COPY --from=builder /go/src/kutala/controller /controller
COPY --from=builder /go/src/kutala/injection /injection
COPY --from=builder /go/src/kutala/utils /utils

EXPOSE 9999

CMD [ "./main" ]