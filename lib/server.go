package lib

import (
	"github.com/razshare/frizzante"
	"main/lib/notifiers"
)

var Server = frizzante.
	NewServer().
	WithNotifier(notifiers.Console).
	WithAddress("127.0.0.1:8080")
