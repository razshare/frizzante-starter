package main

import (
	"embed"
	"main/lib"
	"main/lib/routes"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	routes.Load()
	lib.Server.Start(dist)
}
