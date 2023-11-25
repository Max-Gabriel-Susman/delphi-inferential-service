
local-restart: local-stop local-start

build:
	docker build --tag delphi-model-service .

run: 
	docker run \
		-e API_ADDRESS=0.0.0.0:8082 \
		-e INFERENTIAL_DB_USER=usr \
		-e INFERENTIAL_DB_PASSWORD=identity \
		-e INFERENTIAL_DB_HOST=127.0.0.1 \
		-e INFERENTIAL_DB_NAME=identity \
		-e INFERENTIAL_DB_PORT=3306 \
		-e ENABLE_MIGRATE=true \
		-p 50054:50054 \
		brometheus/delphi-inferential-service:v0.4.5

push: 
	docker push brometheus/delphi-model-service:tagname

update:
	docker build --tag brometheus/delphi-inferential-service:v0.4.5 .
	docker push brometheus/delphi-inferential-service:v0.4.5
]