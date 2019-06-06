#Version 2.0

#源镜像
FROM golang:1.11.2

#设置工作目录
#WORKDIR /go/src

#作者
MAINTAINER xiaofei "xiaofei.liu@wenba100.com"

#将服务器的go工程代码加入到docker容器中
ADD . /go/src

#go构建可执行文件

RUN cd /go/src &&\
	rm Dockerfile &&\
	go install main &&\
	cd /go/bin &&\
	cp main service_shop &&\
	rm main &&\
	#rm -rf /go/src/main &&\
	mkdir /data/ &&\
	touch /data/shop.log &&\
	export GOPATH="/go"
 

#暴露端口
EXPOSE 8000
EXPOSE 8080

#最终运行docker的命令
#CMD ["/go/bin/service_shop > /data/shop.log 2>&1"]
CMD ["/bin/sh", "-c", "/go/bin/service_shop > /data/shop.log 2>&1"]
