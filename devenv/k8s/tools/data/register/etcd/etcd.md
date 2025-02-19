etcdctl --endpoints=https://127.0.0.1:2379 endpoint status --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key

etcd --advertise-client-urls=https://192.168.49.2:2379 \
    --cert-file=/etc/kubernetes/pki/etcd/server.crt \
    --data-dir=/var/lib/minikube/etcd \
    --initial-advertise-peer-urls=https://192.168.49.2:2380 \
    --initial-cluster=minikube=https://192.168.49.2:2380 \
    --key-file=/etc/kubernetes/pki/etcd/server.key \
    --listen-client-urls=https://127.0.0.1:2379,https://192.168.49.2:2379 \
    --listen-metrics-urls=http://127.0.0.1:2381 \
    --listen-peer-urls=https://192.168.49.2:2380 \
    --name=minikube \
    --peer-cert-file=/etc/kubernetes/pki/etcd/peer.crt \
    --peer-key-file=/etc/kubernetes/pki/etcd/peer.key \
    --peer-trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt \
    --proxy-refresh-interval=70000 \
    --snapshot-count=10000 \
    --trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt \
    
etcdctl --endpoints=http://192.168.1.212:2379 endpoint status 