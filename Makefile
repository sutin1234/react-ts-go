server-dev:
	cd go-server && air -c .air.toml

server-build:
	cd go-server && go build -o bin/main

client-dev:
	npm run dev

client-build:
	npm run build