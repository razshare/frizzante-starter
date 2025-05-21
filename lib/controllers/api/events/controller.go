package events

import (
	"fmt"
	f "github.com/razshare/frizzante"
	"time"
)

var Controller = f.
	NewApiController().
	WithPath("/api/events").
	WithHandler("GET", get)

func get(req *f.Request, res *f.Response) {
	alive := req.IsAlive()
	res.SendSseUpgrade()
	for *alive {
		now := time.Now().Format(time.TimeOnly)
		message := fmt.Sprintf("Server time is %s", now)
		res.SendMessage(message)
		time.Sleep(time.Second)
	}
}
