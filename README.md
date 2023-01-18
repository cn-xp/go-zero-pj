# go-zero-pj
go-zero project

# 准备工作
1. 启动etcd服务
2. 启动redis服务
3. 启动数据库
4. 启动rpc：go run transform.go -f etc/transform.yaml
5. 启动api：go run shorturl.go -f etc/shorturl-api.yaml
6. curl -i "http://localhost:8888/shorten?url=http://www.xiaoheiban.cn" 测试shorten
7. curl -i "http://localhost:8888/expand?shorten=f35b2a" 测试expand

# 性能测试
go test -bench=. -benchmem

# goctl
## 通过url生成model
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gozero" -table="*" -dir=.

## 通过sql文件生成
goctl model mysql ddl -c -src shorturl2.sql -dir .
