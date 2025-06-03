dev:
	make update
	mkdir .dist/server -p && touch .dist/server/server.js
	which bin/air || curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s
	DEV=1 CGO_ENABLED=1 ./bin/air \
	--build.cmd "go build -o bin/app ." \
	--build.bin "bin/app" \
	--build.exclude_dir "out,bin,.sessions,.archive,.frz,.dist,.git,.github,node_modules" \
	--build.exclude_regex "_test.go" \
	--build.include_ext "go" \
	--build.log "go-build-errors.log" & \
	DEV=1 go run lib/tools/aot/main.go
	DEV=1 bunx vite build --watch --ssr .frz/router/server.ts --outDir .dist/server & \
	DEV=1 bunx vite build --watch --outDir .dist/client & \
	wait

build:
	make update
	make generate
	CGO_ENABLED=1 go build -o bin/app .

test:
	make update
	make generate
	CGO_ENABLED=1 go test ./...

generate:
	go run lib/tools/aot/main.go
	bunx vite build --ssr .frz/router/server.ts --outDir .dist/server --emptyOutDir
	bunx vite build --outDir .dist/client --emptyOutDir

clean:
	go clean
	rm bin/app -fr
	rm node_modules -fr
	rm .dist -fr
	rm .frz -fr

update:
	go mod tidy
	bun update

hooks:
	printf "#!/usr/bin/env bash\n" > .git/hooks/pre-commit
	printf "make test" >> .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit