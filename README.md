# k8s上部署golang api服务

#### k8s部署
略

#### 简单go程序编写
```
代码省略100行....
```
API：http://localhost:6767/accounts/1

#### go程序编译
由于在centos上创建镜像，因此使用交叉编译
```
GOOS=linux GOARCH=amd64 go build main.go
```

#### 构建镜像

##### Dockerfile文件
```
# 创建一个空镜像，golang镜像太大
FROM scratch

# 将go程序放到/目录
ADD main /

# 设置 PORT 环境变量
ENV PORT 6767

# 给主机暴露 6767 端口，这样外部网络可以访问你的应用
EXPOSE 6767

# 告诉 Docker 启动容器运行的命令
CMD ["/main"]

```
注意： 由于golang的镜像太大，而本例使用编译好程序构建，因此使用空镜像来构建

##### 构建镜像
将dockerfile和main文件拷贝到centos虚拟机中，执行如下命令构建镜像：
```
# 构建
sudo docker build -t accountservice .
# 运行
sudo docker run -it -p 6767:6767 accountservice
```

##### 发布镜像到docker hub
将镜像推送到docker hub中需要打tag  
```
sudo docker tag accountservice:latest kenfo/test
sudo docker push kenfo/test:latest
```

#### k8s部署

##### 需要的相关yaml文件
1. namespace.yaml
```
apiVersion: v1
kind: Namespace
metadata:
   name: kenfo-test
   labels:
     name: kenfo-test
     istio-injection: enabled
```
2. accountservice-service.yaml
```
apiVersion: v1
kind: Service
metadata:
    name: accountservice
    namespace: kenfo-test
spec:
    selector:
        app: accountservice
    ports:
      - port: 6767
        targetPort: 6767
        protocol: TCP
    type: NodePort
```
3. accountservice-deployment.yaml
```
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
    name: accountservice
    namespace: kenfo-test
    labels:
        app: accountservice
        version: v1
spec:
    replicas: 1
    template: 
        metadata: 
            labels: 
                app: accountservice
        spec:
            containers:
              - name: accountservice
                image: kenfo/test:latest
                imagePullPolicy: Always
                ports:
                - containerPort: 6767
                

```

#### 部署到k8s中
```
kubectl create -f namespace.yaml
kubectl create -f accountservice-service.yaml
kubectl create -f accountservice-deployment.yaml
```

#### 验证部署结果
```
查看accoutservice服务对应的nodeport端口
kubectl get svc -n kenfo-test
```

