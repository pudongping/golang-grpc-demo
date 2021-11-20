# 使用 golang 搭建 grpc 服务

- 下载源码

```shell
git clone https://github.com/pudongping/golang-grpc-demo.git
```

- 下载相关依赖

```shell
go mod tidy
```

## 测试

```shell
# 开启服务端
go run server.go

# 开启客户端
go run client.go
```

你将看到以下内容

服务端控制台输出内容为：

```shell

2021/11/21 07:04:06 grpc server is start 0.0.0.0:8081
2021/11/21 07:31:30 receive user index request: page 2 page_size 5
2021/11/21 07:31:30 receive user view request: uid 1
2021/11/21 07:31:30 receive user post request: name Alex, password 123456, age 27
2021/11/21 07:31:30 receive user delete request: uid 2

```

客户端控制台输出内容为：

```shell

2021/11/21 07:31:30 user index success: Success
list ==> name ==>  Alex  age ==>  26
list ==> name ==>  Harry  age ==>  18
2021/11/21 07:31:30 user view success: Success
view ==> name ==>  Alex  age ==>  26
2021/11/21 07:31:30 user post success: Success
2021/11/21 07:31:30 user delete success: Success

```