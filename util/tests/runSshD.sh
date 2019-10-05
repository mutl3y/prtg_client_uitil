#!/usr/bin/env bash
docker build -t sshd .
docker stop sshd
docker rm sshd
docker create --name=sshd -p 1422:22 --restart unless-stopped sshd
docker start sshd
