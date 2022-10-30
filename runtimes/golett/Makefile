generate:
	@protoc -I api/proto api/proto/*.proto \
		--go_out=./api/proto --go_opt=paths=source_relative \
		--go-grpc_out=./api/proto --go-grpc_opt=paths=source_relative