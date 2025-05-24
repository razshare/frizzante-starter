package main

import (
	"embed"
	"main/lib"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	lib.Server.Start(dist)
}
