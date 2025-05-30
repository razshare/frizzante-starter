package routes

import "github.com/razshare/frizzante"

func GetExpired(req *frizzante.Request, res *frizzante.Response) {
	res.SendView(frizzante.View{
		Name:       "Expired",
		RenderMode: frizzante.RenderModeFull,
		Data:       map[string]string{},
	})
}
