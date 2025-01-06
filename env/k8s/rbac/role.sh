cp /var/lib/minikube/certs/ca.crt  .
cp /var/lib/minikube/certs/ca.key .

# 创建根证书

(umask 077;openssl genrsa -out dev.key 2048)
openssl req -new -key dev.key -out dev.csr -subj "/O=k8s/CN=dev"
openssl  x509 -req -in dev.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out dev.crt -days 365