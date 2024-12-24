
# docker
https://docs.docker.com/engine/install/ubuntu/#installation-methods

## wsl2
vim /etc/wsl.conf
[boot]
systemd=true
wsl.exe --shutdown
```bash
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc(可能耗时很长,可能失败,尝试修改hosts,多次重试)
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
"deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
$(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin 可选(docker-compose-plugin)
sudo gpasswd -a ${USER} docker && newgrp docker
sudo systemctl restart docker
mkdir /etc/docker
vi /etc/docker/daemon.json
{
    "registry-mirrors": ["https://docker.mirrors.ustc.edu.cn"],
    "insecure-registries":["${ip}"],
    "features": {
      "buildkit": true
    }

}

{
    "registry-mirrors": ["https://docker.hoper.xyz"],
    "insecure-registries":["docker.hoper.xyz"]
}
docker login -u 用户名 -p 密码 ${ip}

## ali镜像
wget -qO- https://get.docker.com/ | sh
sudo service docker start
curl -sSL http://acs-public-mirror.oss-cn-hangzhou.aliyuncs.com/docker-engine/internet | sh -
sudo apt-get install linux-image-extra-$(uname -r) linux-image-extra-virtual
sudo apt-get update
sudo apt-get install apt-transport-https ca-certificates
sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
echo "deb https://apt.dockerproject.org/repo ubuntu-xenial main" | sudo tee /etc/apt/sources.list.d/docker.list
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin 可选(docker-compose-plugin)
sudo systemctl enable docker
sudo systemctl start docker
```



