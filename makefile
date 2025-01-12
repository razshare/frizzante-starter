update: go.mod www/package.json
	go mod tidy
	cd www && bun update

clean:
	go clean
	rm cert.pem -f
	rm key.pem -f
	rm out -fr
	rm bin -fr
	rm tmp -fr
	rm www/dist -fr
	mkdir www/dist/server -p
	mkdir www/dist/client -p
	touch www/dist/.gitkeep
	touch www/dist/server/.gitkeep
	touch www/dist/client/.gitkeep

build: www-build main.go  go.mod
	CGO_ENABLED=1 go build -o out/app .

start: www-build main.go  go.mod
	CGO_ENABLED=1 go run main.go

dev: bin go.mod
	DEV=1 CGO_ENABLED=1 ./bin/air \
	--build.cmd "go build -o out/app ." \
	--build.bin "out/app" \
	--build.exclude_dir "out,bin,www" \
	--build.exclude_regex "_text.go" \
	--build.include_ext "go" \
	--build.log "go-build-errors.log" & make www-watch & wait

bin:
	curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s

www-prepare:
	go run prepare/main.go

www-build: www-prepare www/package.json
	make www-build-server & make www-build-client & wait

www-build-server: www/package.json
	cd www && \
	bunx vite build --ssr .frizzante/vite-project/render.server.js --outDir dist/server --emptyOutDir && \
	./node_modules/.bin/esbuild dist/server/render.server.js --bundle --outfile=dist/server/render.server.js --format=esm --allow-overwrite

www-build-client: www/package.json
	cd www && \
	bunx vite build --outDir dist/client --emptyOutDir

www-watch: www-prepare www/package.json
	make www-watch-server & make www-watch-client & wait

www-watch-server: www/package.json
	cd www && \
	bunx vite build --watch --ssr .frizzante/vite-project/render.server.js --outDir dist/server --emptyOutDir && \
	./node_modules/.bin/esbuild dist/server/render.server.js --bundle --outfile=dist/server/render.server.js --format=esm --allow-overwrite

www-watch-client: www/package.json
	cd www && \
	bunx vite build --watch --outDir dist/client --emptyOutDir

test: clean update www-build go.mod
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
