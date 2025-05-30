package handlers

import (
	"fmt"
	"github.com/razshare/frizzante"
	"time"
)

func GetEvents(req *frizzante.Request, res *frizzante.Response) {
	alive := req.IsAlive()
	res.SendSseUpgrade()
	for *alive {
		now := time.Now().Format(time.TimeOnly)
		message := fmt.Sprintf("Server time is %s", now)
		res.SendMessage(message)
		time.Sleep(time.Second)
	}
}
