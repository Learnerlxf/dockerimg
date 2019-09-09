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

#设置工作目录
WORKDIR /data/src

#暴露端口
EXPOSE 8080
EXPOSE 8081
EXPOSE 8082

ENV GOROOT=/usr/local/go \
    PATH="$PATH:/usr/local/go/bin" \
    GOPATH=/data

RUN cd /data/src && \
    git clone https://github.com/Learnerlxf/go-chat.git  && \
    go get github.com/gorilla/websocket && \
    cd /data/src/go-chat/docker && \
    chmod +x ./start.sh && \
    cd /data/src/go-chat && \
    go build -o go-chat main.go && \
    chmod +x go-chat
CMD ["/data/src/go-chat/docker/start.sh"]
