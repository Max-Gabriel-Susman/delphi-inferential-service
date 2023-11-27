
local-restart: local-stop local-start

build:
	docker build --tag delphi-model-service .

run: 
	docker run \
		-e API_KEY \
		-e API_ORG \
		-p 50054:50054 \
		brometheus/delphi-inferential-service:v0.4.5

push: 
	docker push brometheus/delphi-model-service:tagname

update:
	docker build --tag brometheus/delphi-inferential-service:v0.4.5 .
	docker push brometheus/delphi-inferential-service:v0.4.5

# grpcurl -plaintext -v localhost:50054 list Greeter

# grpcurl -plaintext -d '{"name": "tell me Im pretty"}' localhost:50054 Greeter/SayHello
