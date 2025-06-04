dev:
	make update
	make generate
	mkdir dist -p
	touch dist/.gitkeep
	which bin/air || curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s
	DEV=1 CGO_ENABLED=1 ./bin/air \
	--build.cmd "go build -o bin/app ." \
	--build.bin "bin/app" \
	--build.exclude_dir "node_modules,dist,bin,sessions,.archive,.git,.github" \
	--build.exclude_regex "_test.go" \
	--build.include_ext "go" \
	--build.log "go-build-errors.log" & \
	DEV=1 bunx vite build --watch --ssr frz/scripts/server.ts --outDir dist/server & \
	DEV=1 bunx vite build --watch --outDir dist/client & \
	wait

build:
	make update
	make generate
	bunx vite build --ssr frz/scripts/server.ts --outDir dist --emptyOutDir
	bunx vite build --outDir dist/client --emptyOutDir
	CGO_ENABLED=1 go build -o bin/app .

test:
	make update
	make generate
	bunx vite build --ssr frz/scripts/server.ts --outDir dist --emptyOutDir
	bunx vite build --outDir dist/client --emptyOutDir
	CGO_ENABLED=1 go test ./...

generate:
	go run cli/main.go -generate -utilities -out="frz"

update:
	go mod tidy
	bun update

clean:
	go clean
	rm bin/app -fr
	rm dist -fr
	rm sessions -fr
	rm node_modules -fr

hooks:
	printf "#!/usr/bin/env bash\n" > .git/hooks/pre-commit
	printf "make test" >> .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit