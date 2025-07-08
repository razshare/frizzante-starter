########################
###### Composites ######
########################
test: package
	CGO_ENABLED=1 go test

build: package
	CGO_ENABLED=1 go build -o .gen/bin/app .

package-watch: touch
	cd app && \
	bunx vite build --logLevel info --ssr frizzante/scripts/server.ts --outDir dist --emptyOutDir false --watch & \
	cd app && \
	bunx vite build --logLevel info --outDir dist/client --emptyOutDir false --watch & \
	cd app && \
	bunx svelte-check --tsconfig ./tsconfig.json --watch --preserveWatchOutput & \
	wait

package: check touch
	cd app && \
	bunx vite build --logLevel info --ssr frizzante/scripts/server.ts --outDir dist && \
	bunx vite build --logLevel info --outDir dist/client && \
	node_modules/.bin/esbuild dist/server.js --bundle --outfile=dist/server.js --format=cjs --allow-overwrite && \
	touch dist/.gitkeep

check: touch
	cd app && \
	bunx eslint . && \
	bunx svelte-check --tsconfig ./tsconfig.json

########################
###### Primitives ######
########################
dev:
	mkdir .gen/tmp -p
	mkdir app/dist -p
	touch app/dist/.gitkeep
	touch app/dist/server.js
	DEV=1 CGO_ENABLED=1 air & \
	make package-watch & \
	wait

clean:
	go clean
	rm app/dist -fr
	rm .gen/tmp -fr
	rm .vite -fr

touch:
	mkdir app/dist -p
	touch app/dist/.gitkeep
	touch app/dist/server.js

format:
	cd app && \
	bunx prettier --write .

install:
	go mod tidy
	cd app && \
	bun install

update:
	go get -u ./...
	go mod tidy
	cd app && \
	bun update


hooks:
	printf "#!/usr/bin/env bash\n" > .git/hooks/pre-commit
	printf "make test" >> .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit