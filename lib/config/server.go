package config

import (
	"embed"
	f "github.com/razshare/frizzante"
)

//go:embed .dist/*/**
var dist embed.FS
var Server = f.
	NewServer().
	WithEfs(dist).
	WithAddress("127.0.0.1:8080")
