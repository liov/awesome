# postgres迁移

## postgres 表移动到另一个库
pg_dump -t table_to_copy source_db | psql target_db

pg_dump -U postgres -d test | psql -d hoper -U postgres
## 备份恢复
pg_dump -U postgres -p 5432 -d test -f /home/postgres/test12.sql
psql -d test -U postgres -f test12.sql


postgres进行迁移可以使用psql，也可以使用postgres自带工具pg_dump和pg_restore.

命令：

- 备份

pg_dump -h 13.xx.xx.76 -U postgres -n "public" "schema" -f ./schema_backup.gz -Z 9

-h host，备份目标数据库的ip地址

-U 用户名（输入命令后会要求输入密码，也可以使用-w输入密码）

-n 需要导出的schema名称

-f 导出存储的文件

-Z 进行压缩（一般导出文件会占用很大的存储空间，直接进行压缩）

- 恢复

gunzip schema_backup.gz ./ （对导出的压缩文件解压）

psql -U postgres -f ./schema_backup >>restore.log

参数意义与导出一样

坑与tips：

版本，pg_dump的版本要高于目标备份数据库的版本（比如目标数据库是10.3， pg_dump要使用10.3或者10.4）

-Z 是pg_dump提供的压缩参数，默认使用的是gzip的格式，目标文件导出后，可以使用gunzip解压（注意扩展名，有时习惯性命名为.dump 或者.zip，使用gunzip时会报错，要改为.gz）

也可以针对指定的表进行导出操作：

pg_dump -h localhost -U postgres -c -E UTF8 --inserts -t public.t_* > t_taste.sql

--inserts 导出的数据使用insert语句

-c 附带创建表命令

## 比较骚，只适用同版本
1.操作位置：迁移数据库源（旧数据库主机）

找到PostgreSql 的data目录   关闭数据库进程

打包 tar -zcvf pgdatabak.tar.gz data/

------------------------------------------------------------------

2.通过winScp 或者 CRT 等工具拷贝到    迁移目标源（新主机--需安装postgresql）  同样的data目录 关闭数据库进程

解压  tar -zxvf pgdatabak.tar.gz -C /usr/local/postgres/

重新授权 执行命令  chown -R postgres:postgres data/

pg_dumpall -U postgres -p 5432 > bak.sql
psql -U postgres -f bak.sql

kubectl exec pod_name  -n tools -- pg_dumpall -U postgres -p 5432 > bak.sql
kubectl exec pod_name  -n tools -- pg_dump -U postgres -p 5432 -d test > bak.sql


kubectl exec -i pod_name  -n tools -- psql -U postgres < bak.sql
--inserts #insert语句导出
kubectl exec postgres-old --  pg_dumpall -U postgres | kubectl exec -i -- postgres-new psql -U postgres

# 备份指定数据范围

CREATE TABLE temp_table AS SELECT * FROM your_table WHERE created_at >= '2025-01-06';
kubectl exec pod_name  -n tools -- pg_dump -U postgres -p 5432 -d dbname -t temp_table > bak.sql
DROP TABLE temp_table;

kubectl exec pod_name  -n tools -- psql -U postgres -p 5432 -d dbname > bak.sql
INSERT INTO your_table SELECT * FROM temp_table ON CONFLICT (unique_column) DO NOTHING;
DROP TABLE temp_table;

## -n参数是无效的
坑逼,-t schema.table
# 设置时区
kubectl exec -it pod_name  -n tools --  psql -U postgres;

set time zone "Asia/Shanghai";
SET TIMEZONE='Asia/Shanghai';

# 改配置
vim postgresql.conf
log_timezone = 'Asia/Shanghai'
timezone = 'Asia/Shanghai'

vim /home/postgres/data/pg_hba.conf
host    all     all     0.0.0.0/0        md5
host    all     all     ::/0             md5