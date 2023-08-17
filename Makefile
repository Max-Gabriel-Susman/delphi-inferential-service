
# protoc --go_out=. --go_opt=paths=source_relative \
# 	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
# 	helloworld/helloworld.proto
gen_infer_proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/protos/inference/inference.proto
	
# protoc -I=proto --go_out=plugins=grpc:proto proto/*.proto

echo:
	python3 internal/protos/inference/test.py