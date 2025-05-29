package sessions

import (
	"github.com/razshare/frizzante"
	"main/lib/notifiers"
	"time"
)

var Archive = frizzante.
	NewArchiveOnDisk(".sessions", time.Second/2).
	WithNotifier(notifiers.Console)
