Rider
vs
Clion
RustRover

# vscode远程开发
```sh
sudo useradd -s /bin/bash -G docker,root -m newuser
sudo passwd newuser
#sudo usermod -aG sudo newuser
#sudo usermod -s /bin/bash newuser
```

ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
sudo mkdir -p /home/newuser/.ssh
sudo chmod 700 /home/newuser/.ssh
cat /tmp/id_rsa.pub >> /home/newuser/.ssh/authorized_keys
sudo chmod 600 /home/newuser/.ssh/authorized_keys
sudo chown -R newuser:newuser /home/newuser/.ssh
rm /tmp/id_rsa.pub
sz id_rsa

本地创建ssh/config
Host $host
  HostName $host
  Port 22



  User newuser
  ForwardAgent yes

#  Permissions for 'id_rsa' are too open.
> It is required that your private key files are NOT accessible by others.
> This private key will be ignored.
> Load key "id_rsa": bad permissions

手动修改文件权限
zsh
chmod 400 id_rsa(无效)


这个时候我们需要添加一个用户，给予访问权限，要与使用 ssh 连接登录的用户一致
删除其他用户,如果无法删除,高级那里点击禁用继承
右键 -》属性 -》 安全 -》 高级 -》 添加 -》 选择主体 -》 高级 -》 立即查找 -》 选择用户后确认，一路保存即可
