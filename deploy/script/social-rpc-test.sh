#!/bin/bash
reso_addr='crpi-641ut6s2k620a3al.cn-beijing.personal.cr.aliyuncs.com/chat-mic/social-rpc-dev'
tag='latest'

pod_ip="8.133.207.187"

container_name="easy-chat-social-rpc-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network imooc_easy-im -v /easy-im/config/user-rpc:/user/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 10001:10001 -e POD_IP=${pod_ip}  --name=${container_name} -d ${reso_addr}:${tag}