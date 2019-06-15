## 设计目标

token-server 每秒产生固定数量的 token，没接收到一个 http 请求就释放一个 token。如果"本秒"产生的 token 被消费完，其他的请求就只能等待下一秒才能返回。另外如果上一秒的 token 没有被消费完，那么下一秒只是补全 token 的总量，不会叠加生成 token。
token-server 的目的是模拟一个标准的 http-server 的行为。同时为了规避依赖服务带来的影响，只是强制模拟成一个只能提供固定 qps 服务的 server。

## Knative Service 

Knative Service 定义如下

```
apiVersion: serving.knative.dev/v1alpha1
kind: Service
metadata:
  name: token-server
  namespace: default
spec:
  template:
    metadata:
      labels:
        app: token-server
      annotations:
        autoscaling.knative.dev/maxScale: "20"
        autoscaling.knative.dev/target: "3"
    spec:
      containers:
        - image: registry.cn-hangzhou.aliyuncs.com/knative-sample/token-server:20190615121805
          ports:
            - name: http1
              containerPort: 8080
          env:
            - name: MESSAGE
              value: "token-server-msg"
            - name: RATE
              value: "6"
```

参数说明
- 环境变量 RATE

表示每秒生成几个 token

- 环境变量 MESSAGE 

表示设定的一个 message 值，用于测试使用。

## 编译镜像

`./build-image.sh`

## HTTP 请求 

其中 sleep=400 表示服务端接收到请求以后 sleep 多长时间，这个用户模拟真实的 RT 
```
curl http://token-server.default.example.com?sleep=400 
```
