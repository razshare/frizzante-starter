test:
	make update
	make generate
	make check
	make package
	CGO_ENABLED=1 go test ./...

build:
	make update
	make generate
	make check
	make package
	CGO_ENABLED=1 go build -o bin/app .

dev:
	make update
	make generate
	make check
	mkdir app/dist/client -p
	touch app/dist/client/index.html
	which bin/air || curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s
	DEV=1 CGO_ENABLED=1 ./bin/air \
	--build.cmd "make package && go build -o bin/app ." \
	--build.bin "bin/app" \
	--build.exclude_dir "node_modules,app/dist,bin,sessions,.archive,.git,.github" \
	--build.exclude_regex "_test.go" \
	--build.include_ext "go,svelte,js,json,ts,html" \
	--build.log "go-build-errors.log" & \
	wait

format:
	bunx prettier --write .

clean:
	go clean
	rm bin/app -fr
	rm app/dist -fr
	rm app/lib/utilities -fr
	rm node_modules -fr

update:
	go mod tidy
	bun update

check:
	bunx eslint .
	bunx svelte-check --tsconfig ./tsconfig.json

generate:
	go run cli/main.go -generate -utilities -out="app/lib/utilities"

package:
	bunx vite build --logLevel info --ssr app/lib/utilities/scripts/server.ts --outDir app/dist --emptyOutDir
	bunx vite build --logLevel info --outDir app/dist/client --emptyOutDir

hooks:
	printf "#!/usr/bin/env bash\n" > .git/hooks/pre-commit
	printf "make test" >> .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit