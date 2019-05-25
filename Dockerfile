FROM golang:latest as builder
WORKDIR go/src/github.com/glefloch/zbootcamp

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o zbootcamp

FROM scratch
COPY --from=builder /go/src/github.com/glefloch/zbootcamp/zbootcamp /zbootcamp

ENTRYPOINT ["/zbootcamp"]
