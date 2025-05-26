package routes

import "github.com/razshare/frizzante"

func GetAny(req *frizzante.Request, res *frizzante.Response) {
	res.SendFileOrElse(func() {
		GetWelcome(req, res)
	})
}
