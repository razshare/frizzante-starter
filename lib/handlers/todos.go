package handlers

import (
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libsession"
	"github.com/razshare/frizzante/libview"
	"main/lib"
)

func Todos(c *libcon.Connection) {
	state, _ := libsession.Session(c, lib.NewState())
	c.SendView(libview.View{Name: "Todos", Data: map[string]any{
		"todos": state.Todos,
	}})
}
