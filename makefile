test:
	make update
	make check
	rm app/dist -fr
	mkdir app/dist/client -p
	touch app/dist/client/index.html
	make package
	CGO_ENABLED=1 go test ./...

build:
	make update
	make check
	rm app/dist -fr
	mkdir app/dist/client -p
	touch app/dist/client/index.html
	make package
	CGO_ENABLED=1 go build -o bin/app .

dev:
	make update
	make check
	mkdir app/dist/client -p
	touch app/dist/client/index.html
	DEV=1 CGO_ENABLED=1 ./bin/air \
	--build.cmd "make package && go build -o bin/app ." \
	--build.bin "bin/app" \
	--build.exclude_dir "node_modules,app/dist,bin,sessions,.archive,.git,.github" \
	--build.exclude_regex "_test.go" \
	--build.include_ext "go,svelte,js,json,ts,html" \
	--build.log "go-build-errors.log" & \
	wait

format:
	cd app && \
	../bin/bun x prettier --write .

clean:
	go clean
	rm app/dist -fr
	mkdir app/dist/client -p
	touch app/dist/client/index.html
	rm app/node_modules -fr

update:
	go mod tidy
	cd app && \
	../bin/bun update

check:
	cd app && \
	../bin/bun x eslint . && \
	../bin/bun x svelte-check --tsconfig ./tsconfig.json

package:
	cd app && \
	../bin/bun x vite build --logLevel info --ssr lib/utilities/scripts/server.ts --outDir dist --emptyOutDir && \
	../bin/bun x vite build --logLevel info --outDir dist/client --emptyOutDir
	app/node_modules/.bin/esbuild app/dist/server.js --bundle --outfile=app/dist/server.js --format=cjs --allow-overwrite

configure:
	mkdir bin -p
	which bin/bun || \
	(curl -fsSL https://github.com/oven-sh/bun/releases/latest/download/bun-linux-x64.zip -o bin/bun.zip && \
	unzip -j bin/bun.zip -d bin && rm bin/bun.zip -f)
	which bin/air || curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s
	go run cli/main.go -generate -utilities -out="app/lib/utilities"

hooks:
	printf "#!/usr/bin/env bash\n" > .git/hooks/pre-commit
	printf "make test" >> .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit