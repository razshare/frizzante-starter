dev:
	which bin/air || curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s
	DEV=1 CGO_ENABLED=1 ./bin/air \
	--build.cmd "go build -o bin/app ." \
	--build.bin "bin/app" \
	--build.exclude_dir "out,bin,.sessions,.archive,.generated,.dist,.git,.github,node_modules" \
	--build.exclude_regex "_test.go" \
	--build.include_ext "go" \
	--build.log "go-build-errors.log" & \
	DEV=1 bunx vite build --watch --outDir .dist/client & \
	DEV=1 bunx vite build --watch --ssr .generated/render/server.ts --outDir .dist/server & \
	wait

build:
	bunx vite build --outDir .dist/client --emptyOutDir
	bunx vite build --ssr .generated/render/server.ts --outDir .dist/server --emptyOutDir
	./node_modules/.bin/esbuild .dist/server/server.js --bundle --outfile=.dist/server/server.js --format=esm --allow-overwrite
	CGO_ENABLED=1 go build -o bin/app .

test:
	CGO_ENABLED=1 go test ./...

generate:
	go run cli/main.go -generate -render -views="lib/components/views" -out=".generated/render"
	go run cli/main.go -generate -utilities -out=".generated/utilities"

update:
	go mod tidy
	bun update

clean:
	go clean
	rm bin/app -fr
	rm node_modules -fr
	rm .sessions -fr
	rm .dist -fr
	rm .generated -fr

hooks:
	printf "#!/usr/bin/env bash\n" > .git/hooks/pre-commit
	printf "make test" >> .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit