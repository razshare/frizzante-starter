package config

import f "github.com/razshare/frizzante"

var Server = f.
	NewServer().
	WithAddress("127.0.0.1:8080")
