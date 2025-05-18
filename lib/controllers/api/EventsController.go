package api

import (
	"fmt"
	f "github.com/razshare/frizzante"
	"time"
)

type EventsController struct {
	f.ApiController
}

func (_ EventsController) Configure() f.ApiConfiguration {
	return f.ApiConfiguration{
		Pattern: "GET /api/events",
	}
}

func (_ EventsController) Handle(request *f.Request, response *f.Response) {
	alive := request.IsAlive()
	response.SendSseUpgrade()

	for *alive {
		now := time.Now().Format(time.TimeOnly)
		message := fmt.Sprintf("Server time is %s", now)
		response.SendMessage(message)
		time.Sleep(time.Second)
	}
}
