FROM golang:latest as builder
WORKDIR /go/src/github.com/glefloch/zbootcamp

COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o zbootcamp


FROM scratch
COPY --from=builder /go/src/github.com/glefloch/zbootcamp/zbootcamp /zbootcamp

ENTRYPOINT ["/zbootcamp"]
