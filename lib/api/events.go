package api

import (
	"fmt"
	f "github.com/razshare/frizzante"
	"time"
)

func RequestIsAlive(req *f.Request) *bool {
	value := true
	go func() {
		<-f.ReceiveCancellation(req)
		value = false
	}()
	return &value
}

func Events(
	withPattern func(pattern string),
	withHandler func(handler func(req *f.Request, res *f.Response)),
) {
	withPattern("GET /api/events")
	withHandler(func(req *f.Request, res *f.Response) {
		alive := RequestIsAlive(req)
		f.SendSseUpgrade(res)

		for *alive {
			f.SendEcho(res, fmt.Sprintf("Server time is %s", time.Now().Format(time.RFC1123)))
			time.Sleep(time.Second)
		}
	})
}
