package handlers

import "github.com/razshare/frizzante"

func GetExpired(req *frizzante.Request, res *frizzante.Response) {
	res.SendView(frizzante.View{Name: "Expired"})
}
