.PHONY: backend-gen-proto
backend-gen-proto:
	docker compose exec backend sh -c "protoc -I /proto --go_out=. --go-grpc_out=. /proto/*.proto"

.PHONY: frontend-gen-proto
frontend-gen-proto:
	docker compose exec frontend sh -c "\
	rm -rf /app/src/grpc/generated && mkdir -p /app/src/grpc/generated; \
	/app/node_modules/.bin/grpc_tools_node_protoc \
	--plugin=protoc-gen-ts=/app/node_modules/.bin/protoc-gen-ts \
	--js_out=import_style=commonjs,binary:/app/src/grpc/generated \
	--grpc_out=grpc_js:/app/src/grpc/generated \
	--ts_out=service=grpc-node,mode=grpc-js:/app/src/grpc/generated \
	-I /proto \
	/proto/*.proto\
	"

.PHONY: backend-go-mod-tidy
backend-go-mod-tidy:
	docker-compose exec backend sh -c "go mod tidy"
