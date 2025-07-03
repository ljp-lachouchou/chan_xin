#!/bin/bash
# 修复shebang

# 正确定义数组：仅包含有效脚本路径
need_start_server_shell=(
  "user-rpc-test.sh" # rpc启动
  "user-api-test.sh"
  "social-rpc-test.sh"
  "social-api-test.sh"
  "im-ws-test.sh"
  "im-rpc-test.sh"
  "im-api-test.sh"
  "task-mq-test.sh"
  "dynamics-rpc-test.sh"
  "dynamics-api-test.sh"
)

for i in "${need_start_server_shell[@]}"; do  # [1](@ref)
  [[ -f "$i" ]] || { echo "⚠️ 文件不存在: $i"; continue; }
  chmod +x "$i"
  ./"$i"
done

docker ps

docker exec -it etcd etcdctl get --prefix ""