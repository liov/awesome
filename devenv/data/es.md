命令行连接 Elasticsearch 的几种方法
以下是使用命令行工具连接 Elasticsearch 的几种常用方式：

# 使用 curl命令
```bash

# 基本连接测试
curl -X GET "http://localhost:9200/"

# 带认证的连接（用户名密码）
curl -u username:password -X GET "http://localhost:9200/"

# 查询集群健康状态
curl -X GET "http://localhost:9200/_cluster/health?pretty"

# 查询索引列表
curl -X GET "http://localhost:9200/_cat/indices?v"
```
# 使用 elasticsearch-cli(官方工具)
```bash
# 安装
npm install -g elasticsearch-cli

# 连接
escli --host http://localhost:9200

# 或者带认证
escli --host http://localhost:9200 --username elastic --password yourpassword
```
# 使用 elasticsearch-sql-cli
```bash
# 安装
pip install elasticsearch-sql-cli

# 连接
elasticsearch-sql-cli http://localhost:9200

# 执行SQL查询
SELECT * FROM index_name LIMIT 10;
```
# 使用 jq处理 JSON 输出
```bash
   curl -s "http://localhost:9200/_cat/nodes?v" | jq
```
# 对于启用了 HTTPS 和认证的集群：
```bash
# 使用证书
curl --cacert /path/to/ca.pem -u elastic:password https://localhost:9200

# 跳过证书验证（仅测试环境）
curl -k -u elastic:password https://localhost:9200
```
# 常用 API 示例
```bash
# 创建索引
curl -X PUT "http://localhost:9200/my_index"

# 添加文档
curl -X POST "http://localhost:9200/my_index/_doc" -H 'Content-Type: application/json' -d'
{
  "title": "Test Document",
  "content": "This is a test document."
}
'

# 搜索文档
curl -X GET "http://localhost:9200/my_index/_search?q=title:test"
```
# 查询 Elasticsearch 索引前 10 条数据的几种方法
## 使用 _searchAPI
```bash
# 基本查询（返回前10条）
curl -X GET "http://localhost:9200/room_new/_search?pretty"

# 明确指定size参数
curl -X GET "http://localhost:9200/your_index_name/_search?size=10&pretty"
```
## 使用 _searchAPI 带请求体
```bash
curl -X GET "http://localhost:9200/your_index_name/_search" -H 'Content-Type: application/json' -d'
{
"size": 10,
"query": {
"match_all": {}
}
}
'
```
## 使用 _catAPI（简洁格式）
```bash
# 只显示ID和内容
curl -X GET "http://localhost:9200/_cat/indices/your_index_name?v&h=index,docs.count"

# 查看文档（显示前10条）
curl -X GET "http://localhost:9200/your_index_name/_search?size=10&pretty=true&filter_path=hits.hits"
```
## 使用 elasticsearch-sql-cli
```bash
# 安装后执行
elasticsearch-sql-cli http://localhost:9200

# 在交互界面中输入
SELECT * FROM your_index_name LIMIT 10;
```

## 带认证的查询
```bash
# 基础认证
curl -u username:password -X GET "http://localhost:9200/your_index_name/_search?size=10&pretty"

# 使用API Key
curl -H "Authorization: ApiKey YOUR_API_KEY" -X GET "http://localhost:9200/your_index_name/_search?size=10"
```

## 指定返回字段
```bash
curl -X GET "http://localhost:9200/your_index_name/_search" -H 'Content-Type: application/json' -d'
{
  "size": 10,
  "_source": ["field1", "field2"],
  "query": {
    "match_all": {}
  }
}
'
```
# 使用 jq处理结果
```bash
curl -s "http://localhost:9200/your_index_name/_search?size=10" | jq '.hits.hits[]'
```