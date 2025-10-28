监听 0.0.0.0
sudo vim /etc/redis/redis.conf
bind 0.0.0.0
protected-mode no

sudo systemctl restart redis