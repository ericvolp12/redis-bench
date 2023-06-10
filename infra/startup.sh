#!/bin/bash

# Update the instance
sudo apt-get update -y

# Install Docker
sudo apt-get install apt-transport-https ca-certificates curl gnupg lsb-release -y
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo \
  "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
  "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update -y
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y

# Add 'user' to docker group so we can execute docker commands without sudo
sudo usermod -aG docker user

# Install Git
sudo apt-get install git btop make -y

# Install Go
curl -LO https://golang.org/dl/go1.20.5.linux-arm64.tar.gz
tar -C /usr/local -xzf go1.20.5.linux-arm64.tar.gz
rm go1.20.5.linux-arm64.tar.gz

# Download required docker images
docker pull redis/redis-stack-server:latest
docker pull docker.dragonflydb.io/dragonflydb/dragonfly:latest

# Clone the git repo for redis testing
git clone https://github.com/ericvolp12/redis-bench.git /home/user/redis-bench

# Add go to path
echo "export PATH=$PATH:/usr/local/go/bin" >> /home/user/.bashrc

# Chown the repo to user
sudo chown -R user:user /home/user
