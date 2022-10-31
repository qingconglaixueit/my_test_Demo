# go-zero 快速实战 dmeo

## 环境搭建

可以查看文章：https://juejin.cn/post/7138960256054345741

## 代码启动

启动 tenant  rpc 服务：

```
git clone git@github.com:qingconglaixueit/my_test_Demo.git
cd my_test_demo/mymall/tenant/rpc/
go mod tidy
go run tenant.go
```

启动 order api 服务：

```
cd my_test_demo/mymall/order/api/
go run order.go
```

## 效果验证

```
curl -i "http://localhost:9998/api/order/get/约等于公司"
```

```
]# curl -i "http://localhost:9998/api/order/get/约等于公司"
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Traceparent: 00-ebdfab71ccd5c313c59b41c1e168509b-ddd324acb0919c09-00
Date: Sat, 06 Aug 2022 08:24:51 GMT
Content-Length: 63
```
