#!/bin/bash
set -e

# 安装 docker
curl -fsSL https://get.docker.com | sh && \

systemctl start docker && systemctl enable docker && \

echo '{
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "500m",
    "max-file": "20"
  }
}'

> /etc/docker/daemon.json

# 安装 docker-compose
curl -L "https://github.com/docker/compose/releases/download/1.29.2//docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && \

chmod +x /usr/local/bin/docker-compose && \

yum install -y vim git wget tmux bash-completion && \

# 安装 ctop
wget https://github.com/bcicen/ctop/releases/download/0.7.6/ctop-0.7.6-linux-amd64 -O /usr/local/bin/ctop && \

chmod +x /usr/local/bin/ctop