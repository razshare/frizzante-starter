package lib

import "github.com/razshare/frizzante"

var Server = frizzante.
	NewServer().
	WithNotifier(notifier).
	WithAddress("127.0.0.1:8080")
