package lib

import (
	"github.com/razshare/frizzante"
	"main/lib/notifier"
)

var Server = frizzante.
	NewServer().
	WithNotifier(notifier.Console).
	WithAddress("127.0.0.1:8080")
