
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



mod:
	go mod tidy 
	go mod vendor 

local-build-up:
	echo "dockerize this bitch s'il vouz plait"

local-build-down:
	echo "go to sleep"

local-test:
	echo "congrats you're HIV positive!"

# we're gonna wanna parameterize this bad boi later
local-db-up:
	 docker run -d --name ms -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password mysql
# docker exec -d ms bash mysql-u root -ppassword # i don't think I actually need this command since I'm not mannually logging in

local-db-migrate:
	echo "is this even necessary for GORM? or is this only for schema first ORMs like sql-c?"
	 

# okay so it looks like we've managed to login successfully w/ the script
# now we need to find the right connection string 

# this need a check to make sure that mysql is actuall up 
local-db-down:
	docker stop ms 
	docker rm ms


# idk but the mod target was breaking
# local-env-up: mod local-db-up local-db-migrate local-build-up local-test 
local-env-up: local-db-up local-db-migrate local-build-up local-test 
	

local-env-down: local-build-down local-db-down 

local-stop:
	docker kill identity

local-start-with-db:
	docker run --name identity -d \
		--network=bridge --rm \
		-e MYSQL_USER=usr \
		-e MYSQL_PASSWORD=identity \
		-e MYSQL_DATABASE=identity \
		-e MYSQL_ALLOW_EMPTY_PASSWORD=yes \
		-p 3306:3306 mysql

	DD_DISABLE=true \
		API_ADDRESS=0.0.0.0:8080 \
		IDENTITY_DB_USER=usr \
		IDENTITY_DB_PASSWORD=identity \
		IDENTITY_DB_HOST=127.0.0.1 \
		IDENTITY_DB_NAME=identity \
		IDENTITY_DB_PORT=3306 \
		ENABLE_MIGRATE=true \
		go run ./cmd/bestir-identity-service/


local-start:
	DD_DISABLE=true \
		API_ADDRESS=0.0.0.0:8080 \
		INFERENTIAL_DB_USER=usr \
		INFERENTIAL_DB_PASSWORD=identity \
		INFERENTIAL_DB_HOST=127.0.0.1 \
		INFERENTIAL_DB_NAME=identity \
		INFERENTIAL_DB_PORT=3306 \
		ENABLE_MIGRATE=true \
		go run ./cmd/delphi-inferential-service/main.go

local-restart: local-stop local-start