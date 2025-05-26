package routes

import (
	"github.com/razshare/frizzante"
	"main/lib"
	"main/lib/guards"
)

func init() {
	lib.Server.WithRequestHandler("GET /welcome", GetWelcome)
}

func GetWelcome(req *frizzante.Request, res *frizzante.Response) {
	if !frizzante.AllGuardsPass(req, res, guards.NotExpired) {
		return
	}

	res.SendView(frizzante.View{
		Name:       "Welcome",
		RenderMode: frizzante.RenderModeFull,
		Data:       map[string]string{},
	})

}
