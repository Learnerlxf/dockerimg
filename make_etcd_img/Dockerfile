#源镜像
FROM centos:7.5.1804
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
RUN yum -y install rsync && rm -rf /var/cache/yum
RUN yum -y install nc && rm -rf /var/cache/yum
RUN curl -L https://github.com/coreos/etcd/releases/download/v3.3.2/etcd-v3.3.2-linux-amd64.tar.gz -o ./etcd-v3.3.2-linux-amd64.tar.gz \
    && tar xzvf etcd-v3.3.2-linux-amd64.tar.gz \
    && mv etcd-v3.3.2-linux-amd64/{etcd,etcdctl} /usr/local/bin/ \
    && rm -rf etcd-v3.3.2-linux-amd64.tar.gz etcd-v3.3.2-linux-amd64
RUN mkdir -p /go/src/k8s.io/kubernetes \
    && mkdir /data
#设置工作目录
WORKDIR /data

#暴露端口
EXPOSE 2379 2380 3379 3380

ENV GOROOT=/usr/local/go \
    PATH="$PATH:/usr/local/go/bin" \
    GOPATH=/go
ADD run.sh /usr/local/bin/run.sh

#最终运行docker的命令
RUN chmod +x /usr/local/bin/run.sh
#ENTRYPOINT ["/data/run.sh"]
CMD ["/bin/sh", "-c", "run.sh"]
