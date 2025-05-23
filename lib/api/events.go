package api

import (
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib/config"
	"time"
)

var guards []f.Guard

func init() {
	config.Server.OnRequest("GET /api/events", guards, func(req *f.Request, res *f.Response) {
		alive := req.IsAlive()
		res.SendSseUpgrade()
		for *alive {
			now := time.Now().Format(time.TimeOnly)
			message := fmt.Sprintf("Server time is %s", now)
			res.SendMessage(message)
			time.Sleep(time.Second)
		}
	})
}
