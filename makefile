test:
	make update
	make generate
	make package
	CGO_ENABLED=1 go test ./...

build:
	make update
	make generate
	make package
	CGO_ENABLED=1 go build -o bin/app .

dev:
	make update
	make generate
	make package
	DEV=1 bunx vite build --watch --ssr app/lib/utilities/scripts/server.ts --outDir app/dist & \
	DEV=1 bunx vite build --watch --outDir app/dist/client & \
	which bin/air || curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s
	DEV=1 CGO_ENABLED=1 ./bin/air \
	--build.cmd "go build -o bin/app ." \
	--build.bin "bin/app" \
	--build.exclude_dir "node_modules,dist,bin,sessions,.archive,.git,.github" \
	--build.exclude_regex "_test.go" \
	--build.include_ext "go" \
	--build.log "go-build-errors.log" & \
	wait

clean:
	go clean
	rm bin/app -fr
	rm app/dist -fr
	rm app/lib/utilities -fr
	rm node_modules -fr

update:
	go mod tidy
	bun update

generate:
	go run cli/main.go -generate -utilities -out="app/lib/utilities"

package:
	bunx vite build --ssr app/lib/utilities/scripts/server.ts --outDir app/dist --emptyOutDir
	bunx vite build --outDir app/dist/client --emptyOutDir

hooks:
	printf "#!/usr/bin/env bash\n" > .git/hooks/pre-commit
	printf "make test" >> .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit