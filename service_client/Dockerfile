#源镜像
FROM centos:7.5.1804

#作者
MAINTAINER xiaofei "xiaofei.liu@wenba100.com"

RUN yum upgrade -y && rm -rf /var/cache/yum
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
RUN yum -y install gcc automake autoconf libtool make gcc-c++ && rm -rf /var/cache/yum
RUN mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.old \
    && wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo \
    && yum -y install rsync nc htop && rm -rf /var/cache/yum
RUN yum install -y git && rm -rf /var/cache/yum
RUN echo "10.19.150.79 1v1-classroom.xueba10.com" >> /etc/hosts
RUN mkdir -p /data/HelloWorld/src
#设置工作目录
WORKDIR /data/HelloWorld/src

#暴露端口
EXPOSE 8080
EXPOSE 8081

ENV GOROOT=/usr/local/go \
    PATH="$PATH:/usr/local/go/bin" \
    GOPATH=/data/HelloWorld
#将服务器的go工程代码加入到docker容器中
ADD data /data
RUN go build -o nohost-client /data/HelloWorld/src/main.go
#最终运行docker的命令

CMD ["/bin/sh", "-c", "sh /data/HelloWorld/src/start.sh"]
