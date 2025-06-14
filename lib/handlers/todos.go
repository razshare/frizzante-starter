package handlers

import (
	"github.com/razshare/frizzante/frz"
	"main/lib"
)

func Todos(c *frz.Connection) {
	state, _ := frz.Session(c, lib.NewState())
	c.SendView(frz.View{Name: "Todos", Data: map[string]any{
		"todos": state.Todos,
	}})
}
