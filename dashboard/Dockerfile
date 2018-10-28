# 创建一个空镜像，golang镜像太大
FROM scratch

# 在容器内部创建一个目录来存储我们的 web 应用，接着使它成为工作目录
#RUN mkdir -p /go/src/app
#WORKDIR /go/src/app

# 复制服务到容器中
#COPY main /go/src/app
ADD main /
# 设置 PORT 环境变量
ENV PORT 6767

# 给主机暴露 6767 端口，这样外部网络可以访问你的应用
EXPOSE 6767

# 告诉 Docker 启动容器运行的命令
CMD ["/main"]
