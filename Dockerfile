FROM centos:latest


# 安装docker
RUN yum install -y yum-utils device-mapper-persistent-data lvm2 && \
    yum-config-manager --add-repo  https://download.docker.com/linux/centos/docker-ce.repo && \
    yum-config-manager --enable docker-ce-edge && \
    yum install docker-ce -y && \
    mkdir -p /etc/docker && \
    echo -e "{" > /etc/docker/daemon.json && \
    echo -e "\t\"registry-mirrors\":\"[\"https://f8nxo85s.mirror.aliyuncs.com\"]\"" >> /etc/docker/daemon.json && \ 
    echo -e "}" >> /etc/docker/daemon.json && \ 
    systemctl daemon-reload && \
    systemctl restart docker

# 安装gitlabrunner
RUN docker pull gitlab/gitlab-runner && \
    docker run -d --name gitlab-runner --restart always -v /var/run/docker.sock:/var/run/docker.sock -v /etc/gitlab-runner gitlab/gitlab-runner:latest

RUN docker ps
# 工作目录
WORKDIR /app

CMD ["/usr/sbin/init"]