package sessions

import (
	"github.com/razshare/frizzante"
	"time"
)

var Archive = frizzante.NewArchiveOnDisk(".sessions", time.Second/2)
