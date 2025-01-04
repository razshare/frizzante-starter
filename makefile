configure:
	which bun || (curl -fsSL https://bun.sh/install | bash)

load: ui/package.json go.mod
	cd ui && bun update
	go mod tidy

clean:
	go clean
	rm out -fr
	rm ui/.temp -fr
	rm ui/node_modules -fr

start: main.go
	CGO_ENABLED=1 go run main.go

build: main.go
		CGO_ENABLED=1 go build main.go && mkdir out -p && mv main out/frizzante