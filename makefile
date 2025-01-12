update: www/package.json go.mod
	go mod tidy
	go run prepare/main.go
	make redist
	cd www && bun update
	cd www && bunx vite build --ssr .frizzante/vite-project/render.server.js --outDir dist/server --emptyOutDir
	cd www && ./node_modules/.bin/esbuild dist/server/render.server.js --bundle --outfile=dist/server/render.server.js --format=esm --allow-overwrite
	cd www && bunx vite build --outDir dist/client --emptyOutDir

clean:
	go clean
	rm cert.pem -f
	rm key.pem -f
	rm out -fr
	make redist
	mkdir www/dist -p
	touch www/dist/.gitkeep

redist:
	rm www/dist -fr
	mkdir www/dist/server -p
	mkdir www/dist/client -p
	touch www/dist/.gitkeep
	touch www/dist/server/.gitkeep
	touch www/dist/client/.gitkeep

start: main.go
	CGO_ENABLED=1 go run main.go

build: clean update main.go
	CGO_ENABLED=1 go build main.go && mkdir out -p && mv main out/app

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
