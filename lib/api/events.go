package api

import (
	"fmt"
	f "github.com/razshare/frizzante"
	"time"
)

func RequestIsAlive(request *f.Request) *bool {
	value := true
	go func() {
		<-f.ReceiveCancellation(request)
		value = false
	}()
	return &value
}

func Events(context f.ApiContext) {
	// Context.
	withPattern, withHandler := context()

	// Configure.
	withPattern("GET /api/events")
	withHandler(func(request *f.Request, response *f.Response) {
		alive := RequestIsAlive(request)
		f.SendSseUpgrade(response)

		for *alive {
			f.SendEcho(response, fmt.Sprintf("Server time is %s", time.Now().Format(time.TimeOnly)))
			time.Sleep(time.Second)
		}
	})
}
