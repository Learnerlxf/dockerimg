#源镜像
FROM centos:7.5.1804

#设置工作目录
#WORKDIR /go/src

#作者
MAINTAINER xiaofei "xiaofei.liu@wenba100.com"

RUN yum install -y bind-utils && rm -rf /var/cache/yum
RUN yum install -y net-tools && rm -rf /var/cache/yum
RUN yum install -y wget && rm -rf /var/cache/yum
RUN yum install -y vim && rm -rf /var/cache/yum
RUN wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.12.5.linux-amd64.tar.gz \
    && echo -e "\nexport PATH=\$PATH:/usr/local/go/bin" >> ~/.bash_profile \
    && rm -f go1.12.5.linux-amd64.tar.gz
RUN yum install -y telnet && rm -rf /var/cache/yum
RUN yum install -y bind-utils && rm -rf /var/cache/yum
RUN yum install -y ipvsadm && rm -rf /var/cache/yum
RUN yum install -y iptables && rm -rf /var/cache/yum
#暴露端口
EXPOSE 8000
EXPOSE 8082

ENV GOROOT=/usr/local/go \
    PATH="$PATH:/usr/local/go/bin"
#将服务器的go工程代码加入到docker容器中
ADD data /data
RUN go build /data/main.go

#最终运行docker的命令
CMD ["/bin/sh", "-c", "go run /data/main.go"]
