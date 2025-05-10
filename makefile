test: configure
	CGO_ENABLED=1 go test

build: configure
	CGO_ENABLED=1 go build -o bin/app .

start: configure
	CGO_ENABLED=1 go run .

dev: clean update
	go run lib/prepare/main.go
	which bin/air || curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s
	DEV=1 CGO_ENABLED=1 ./bin/air \
	--build.cmd "go build -o bin/app ." \
	--build.bin "bin/app" \
	--build.exclude_dir "out,tmp,bin,archive" \
	--build.exclude_regex "_test.go,.frizzante,node_modules" \
	--build.include_ext "go,svelte,js,css,json" \
	--build.log "go-build-errors.log" & \
	make www-watch-server & \
	make www-watch-client & \
	wait

www-watch-server:
	bunx vite build --watch --ssr .frizzante/vite-project/render.server.js --outDir .dist/server && \
	./node_modules/.bin/esbuild .dist/server/render.server.js --bundle --outfile=.dist/server/render.server.js --format=esm --allow-overwrite

www-watch-client:
	bunx vite build --watch --outDir .dist/client

configure: clean update
	go run lib/prepare/main.go
	make www-build-server & \
	make www-build-client & \
	wait

clean:
	go clean
	rm main.db -f
	rm cert.pem -f
	rm key.pem -f
	rm bin/app -f
	rm tmp -fr
	rm tmp -fr
	rm node_modules -fr
	rm .dist -fr
	rm .frizzante -fr
	mkdir .dist/server -p
	mkdir .dist/client -p
	touch .dist/.gitkeep
	touch .dist/server/.gitkeep
	touch .dist/client/.gitkeep

update:
	go mod tidy
	bun update

www-build-server:
	bunx vite build --ssr .frizzante/vite-project/render.server.js --outDir .dist/server --emptyOutDir && \
	./node_modules/.bin/esbuild .dist/server/render.server.js --bundle --outfile=.dist/server/render.server.js --format=esm --allow-overwrite

www-build-client:
	bunx vite build --outDir .dist/client --emptyOutDir

certificate-interactive:
	openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem

certificate:
	openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem -nodes -subj \
	"/C=XX/ST=Test/L=Test/O=Test/OU=Test/CN=Test"

hooks:
	printf "#!/usr/bin/env bash\n" > .git/hooks/pre-commit
	printf "make test" >> .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit

api:
	go run lib/make/main.go -api

guard:
	go run lib/make/main.go -guard

page:
	go run lib/make/main.go -page
