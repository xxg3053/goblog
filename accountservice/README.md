# docker运行accountservice服务

#### 交叉编译
```
GOOS=linux GOARCH=amd64 go build main.go
```

#### 运行如下命令
```
sudo docker build -t accountservice .
sudo docker run -it -p 6767:6767 accountservice
```

#### 访问
http://localhost:6767/accounts/1