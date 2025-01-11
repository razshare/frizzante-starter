update: www/package.json go.mod
	mkdir www/dist -p
	go mod tidy
	go run main.prepare.go
	cd www && bun update
	cd www && bunx vite build --ssr frizzante/vite-project/render.server.js --outDir dist/server --emptyOutDir
	cd www && ./node_modules/.bin/esbuild dist/server/render.server.js --bundle --outfile=dist/server/render.server.js --format=esm --allow-overwrite
	cd www && bunx vite build --outDir dist/client --emptyOutDir

clean:
	go clean
	rm cert.pem -f
	rm key.pem -f
	rm out -fr
	rm www/frizzante -fr
	rm www/.temp -fr
	rm www/dist/server -fr
	rm www/dist/client -fr
	rm www/node_modules -fr

start: update main.go
	CGO_ENABLED=1 go run main.go

build: clean update main.go
		CGO_ENABLED=1 go build main.go && mkdir out -p && mv server out/app

test:
	go test

certificate-interactive:
	openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem

certificate:
	openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem -nodes -subj \
	"/C=XX/ST=Test/L=Test/O=Test/OU=Test/CN=Test"

hooks:
	printf "#!/usr/bin/env bash\n" > .git/hooks/pre-commit
	printf "make test" >> .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
