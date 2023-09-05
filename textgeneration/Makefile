.PHONY: protos

# from nics
# protos: 
# 	protoc -I protos/ protos/textGenerationService.proto --go-grpc_out=./protos/textGenerationService

protos:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		textGenerationService.proto

pb-gen-tg:
	protoc --go_out=. --go_opt=paths=./pb/textgeneration \
		--go-grpc_out=. --go-grpc_opt=paths=./pb/textgeneration \
		pb/textGenerationService.proto
