configure:
	curl -fsSL https://bun.sh/install | bash

load: ui/package.json
	cd ui && bun update

clean:
	rm firzzante -f
	rm ui/.temp -fr
	rm ui/node_modules -fr

build: main.go
	CGO_ENABLED=1 go build -o frizzante