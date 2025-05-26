package routes

import (
	"github.com/razshare/frizzante"
	"main/lib"
)

func init() {
	lib.Server.WithRequestHandler("GET /expired", GetExpired)
}

func GetExpired(req *frizzante.Request, res *frizzante.Response) {
	res.SendView(frizzante.View{
		Name:       "Expired",
		RenderMode: frizzante.RenderModeFull,
		Data:       map[string]string{},
	})
}
