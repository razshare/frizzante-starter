update: www/package.json go.mod
	cd www && bun update
	cd www && bunx vite build --ssr render.server.js --outDir dist/server
	cd www && ./node_modules/.bin/esbuild dist/server/render.server.js --bundle --outfile=dist/server/render.server.js --format=esm --allow-overwrite
	cd www && bunx vite build --outDir dist/client
	go mod tidy

clean:
	go clean
	rm out -fr
	rm www/.temp -fr
	rm www/node_modules -fr

start: update main.go
	CGO_ENABLED=1 go run main.go

build: update main.go
		CGO_ENABLED=1 go build main.go && mkdir out -p && mv main out/frizzante