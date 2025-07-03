#!/bin/bash
reso_addr='crpi-6tiftpwidsfx5hyi.cn-beijing.personal.cr.aliyuncs.com/chan-xin/im-ws-dev'
tag='latest'

container_name="chan-xin-im-ws-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network chan-xin -v /chan-xin/config/im-ws:/im/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 10086:10086  --name=${container_name} -d ${reso_addr}:${tag}
