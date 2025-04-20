package api

import (
	"fmt"
	f "github.com/razshare/frizzante"
	"time"
)

func IsAlive(req *f.Request) *bool {
	value := true
	go func() {
		<-f.ReceiveCancellation(req)
		value = false
	}()
	return &value
}

func testServeFunction(req *f.Request, res *f.Response) {
	alive := IsAlive(req)
	f.SendSseUpgrade(res)

	for *alive {
		f.SendEcho(res, fmt.Sprintf("Server time is %s", time.Now().Format(time.RFC1123)))
		time.Sleep(time.Second)
	}
}

func Events(
	route func(pattern string),
	serve func(testServeFunction func(req *f.Request, res *f.Response)),
) {
	route("GET /api/events")
	serve(testServeFunction)
}
