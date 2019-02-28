#!/bin/bash
set -e

curl -fsSL https://get.docker.com | sh && \

curl -L "https://github.com/docker/compose/releases/download/1.23.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && \

chmod +x /usr/local/bin/docker-compose && \

systemctl start docker && systemctl enable docker && \

yum install -y vim git wget tmux bash-completion
