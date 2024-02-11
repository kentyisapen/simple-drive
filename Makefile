backend-gen-proto:
	docker-compose exec backend sh -c "protoc -I /proto --go_out=. --go-grpc_out=. /proto/*.proto"
