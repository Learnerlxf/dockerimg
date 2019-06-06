FROM  registry.cn-hangzhou.aliyuncs.com/nbsgcl/golang:centos7.5

RUN mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup \
    && wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
RUN yum install -y java-1.8.0
RUN wget -O /data/jenkins.war http://mirrors.jenkins.io/war-stable/2.138.3/jenkins.war

RUN mkdir /data/jenkins_home

ENV JENKINS_HOME=/data/jenkins_home
EXPOSE 8081 50000
WORKDIR /data
CMD ["/bin/sh", "-c", "java -jar /data/jenkins.war --httpPort=8081"]
