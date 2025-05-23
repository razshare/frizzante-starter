package api

import (
	"fmt"
	"main/lib/config"
	"time"

	f "github.com/razshare/frizzante"
)

func init() {
	var guards []f.Guard
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
