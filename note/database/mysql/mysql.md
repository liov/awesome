# 库同步
mysqldump -h 192.168.35.213 -uroot -p123456 dbname \
--single-transaction \
--set-gtid-purged=OFF \
--routines \
--events \
--triggers | \
mysql -h 127.0.0.1 -uroot -p123456 dbname

# ERROR 1118 (42000) at line 11268: Row size too large (> 8126). Changing some columns to TEXT or BLOB or using ROW_FORMAT=DYNAMIC or ROW_FORMAT=COMPRESSED may help. In current row format, BLOB prefix of 768 bytes is stored inline.

## 编辑MySQL配置文件
sudo vim /etc/mysql/my.cnf

## 添加以下配置
[mysqld]
innodb_strict_mode=0
innodb_file_format=Barracuda # mysql8移除
innodb_file_per_table=1
innodb_default_row_format=DYNAMIC

## 重启MySQL
sudo systemctl restart mysql