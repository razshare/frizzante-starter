package routes

import (
	"github.com/razshare/frizzante"
	"main/lib"
)

func Load() {
	lib.Server.WithRequestHandler("GET /", func(req *frizzante.Request, res *frizzante.Response) {
		res.SendFileOrElse(func() {
			GetWelcome(req, res)
		})
	})
}
