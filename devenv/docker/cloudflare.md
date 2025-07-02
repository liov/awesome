new worker
proxy.js
vim /etc/docker/daemon.json
{
    "registry-mirrors": ["https://yuming"],
    "insecure-registries": ["yuming"]
}