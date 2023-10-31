ifneq (,$(wildcard .env))
	include .env
	export
else
	$(error No .env file found)
endif

migrate: 
	~/go/bin/tern migrate -m migrations \
 		--host     $(DB_HOST) \
		--port     $(DB_PORT) \
   		--database $(DB_NAME) \
    	--user     $(DB_USER) \
    	--password $(DB_PASSWORD) \

build:
	go build -v -o bin/vk2discord ./cmd/main.go

run:
	bin/vk2discord

tests:
	go test -timeout 10s ./...
