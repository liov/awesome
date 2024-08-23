docker build -t jybl/godlv --build-arg BUILD_IMAGE=golang:1.23.0-alpine3.19 --build-arg RUN_IMAGE=alpine:3.19 -f 
Dockerfile .

```dockerfile

ARG BUILD_IMAGE=jybl/protogen
ARG RUN_IMAGE=frolvlad/alpine-glibc
FROM ${BUILD_IMAGE} AS builder

ENV GOPROXY https://goproxy.io,https://goproxy.cn,direct

RUN go install github.com/go-delve/delve/cmd/dlv@latest

FROM ${RUN_IMAGE} AS runtime

COPY --from=builder /go/bin/dlv /bin/

```

```yaml
apiVersion: v1
kind: Service
metadata:
  name: debug-svc
  namespace: default
  labels:
    app: debug-app
spec:
  type: NodePort
  ports:
    - name: http
      port: 80
      protocol: TCP
      nodePort: 10666
    - name: http
      port: 2345
      protocol: TCP
      nodePort: 12345
  selector:
    app: debug-app
```