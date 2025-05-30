package handlers

import "github.com/razshare/frizzante"

func GetWelcome(req *frizzante.Request, res *frizzante.Response) {
	res.SendView(frizzante.View{Name: "Welcome"})
}
