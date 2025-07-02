goctl model mysql ddl -src="./deploy/sql/dynamics.sql" -dir="./apps/dynamics/dynamicsmodels/" -c
goctl rpc protoc ./apps/dynamics/rpc/dynamics.proto --go_out=./apps/dynamics/rpc/ --go-grpc_out=./apps/dynamics/rpc/ --zrpc_out=./apps/dynamics/rpc/
goctl api go -api ./apps/dynamics/api/dynamics.api -dir ./apps/dynamics/api -style gozero