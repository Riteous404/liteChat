#!/bin/bash
reso_addr='registry.cn-hangzhou.aliyuncs.com/liteChat/social-rpc-dev'
tag='latest'

pod_ip="127.0.0.1"

container_name="liteChat-social-rpc-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}

docker run -p 10001:10001 -e POD_IP=${pod_ip}  --name=${container_name} -d ${reso_addr}:${tag}