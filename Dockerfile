FROM registry.cn-hangzhou.aliyuncs.com/knative-sample/golang:1.12 as builder
WORKDIR /go/src/github.com/knative-samples/token-server/
COPY . .
RUN CGO_ENABLED=0 go build -v -o token-server

FROM registry.cn-hangzhou.aliyuncs.com/knative-sample/alpine-sh:3.9
COPY --from=builder /go/src/github.com/knative-samples/token-server/token-server .
ENTRYPOINT ["/token-server"]
