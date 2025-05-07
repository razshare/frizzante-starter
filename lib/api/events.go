package api

import (
	"fmt"
	f "github.com/razshare/frizzante"
	"time"
)

func RequestIsAlive(request *f.Request) *bool {
	value := true
	go func() {
		<-f.RequestReceiveCancellation(request)
		value = false
	}()
	return &value
}

func Events(api *f.Api) {
	f.ApiWithPattern(api, "GET /api/events")
	f.ApiWithHandler(api, func(request *f.Request, response *f.Response) {
		alive := RequestIsAlive(request)
		f.ResponseSendSseUpgrade(response)

		for *alive {
			f.ResponseSendMessage(response, fmt.Sprintf("Server time is %s", time.Now().Format(time.TimeOnly)))
			time.Sleep(time.Second)
		}
	})
}
