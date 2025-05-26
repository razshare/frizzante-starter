package routes

import (
	"fmt"
	"github.com/razshare/frizzante"
	"main/lib"
	"main/lib/guards"
	"time"
)

func init() {
	lib.Server.WithRequestHandler("GET /events", GetEvents)
}

func GetEvents(req *frizzante.Request, res *frizzante.Response) {
	if !frizzante.AllGuardsPass(req, res, guards.NotExpired) {
		return
	}

	alive := req.IsAlive()
	res.SendSseUpgrade()
	for *alive {
		now := time.Now().Format(time.TimeOnly)
		message := fmt.Sprintf("Server time is %s", now)
		res.SendMessage(message)
		time.Sleep(time.Second)
	}

}
