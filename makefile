update: www/package.json go.mod
	cd www && bun update
	cd www && bunx vite build --ssr render.server.js --outDir dist/server
	cd www && ./node_modules/.bin/esbuild dist/server/render.server.js --bundle --outfile=dist/server/render.server.js --format=esm --allow-overwrite
	cd www && bunx vite build --outDir dist/client
	go mod tidy

clean:
	go clean
	rm out -fr
	rm www/.temp -fr
	rm www/dist/server -fr
	rm www/dist/client -fr
	rm www/node_modules -fr

start: update main.go
	CGO_ENABLED=1 go run main.go

build: update main.go
		CGO_ENABLED=1 go build main.go && mkdir out -p && mv main out/app

test:
	go test

certificate:
	openssl genrsa -out server.key 2048
	openssl ecparam -genkey -name secp384r1 -out server.key
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

hooks:
	printf "#!/usr/bin/env bash\n" > .git/hooks/pre-commit
	printf "make test" >> .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit