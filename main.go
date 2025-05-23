package main

import (
	"embed"
	_ "main/lib/api"
	"main/lib/config"
	_ "main/lib/controllers/any"
	_ "main/lib/controllers/expired"
	_ "main/lib/controllers/todos"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	config.Server.WithEfs(dist).Start()
}
