package events

import (
	"fmt"
	f "github.com/razshare/frizzante"
	"time"
)

type Controller struct{}

func (_ Controller) Configure() f.ApiConfiguration {
	return f.ApiConfiguration{
		Pattern: "GET /api/events",
	}
}

func (_ Controller) Handle(req *f.Request, res *f.Response) {
	alive := req.IsAlive()
	res.SendSseUpgrade()
	for *alive {
		now := time.Now().Format(time.TimeOnly)
		message := fmt.Sprintf("Server time is %s", now)
		res.SendMessage(message)
		time.Sleep(time.Second)
	}
}
