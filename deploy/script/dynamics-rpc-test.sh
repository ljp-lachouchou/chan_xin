#!/bin/bash
reso_addr='crpi-6tiftpwidsfx5hyi.cn-beijing.personal.cr.aliyuncs.com/chan-xin/dynamics-rpc-dev'
tag='latest'

container_name="chan-xin-dynamics-rpc-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network imooc_chan-xin -v /chan-xin/config/dynamics-rpc:/dynamics/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 10003:10003  --name=${container_name} -d ${reso_addr}:${tag}
