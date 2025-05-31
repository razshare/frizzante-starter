package main

import (
	"embed"
	"github.com/razshare/frizzante"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	frizzante.NewServer().WithDist(dist).Start()
}
