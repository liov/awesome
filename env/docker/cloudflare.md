new worker
proxy.js
vim /etc/docker/daemon.json
{
    "registry-mirrors": ["https://替换域名"]
   "insecure-registries": ["替换域名"]"
}