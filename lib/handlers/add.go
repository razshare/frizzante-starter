package handlers

import (
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libsession"
	"github.com/razshare/frizzante/libview"
	"main/lib"
)

func Add(c *libcon.Connection) {
	state, operator := libsession.Session(c, lib.NewState())
	defer operator.Save(state)

	description := c.ReceiveQuery("description")

	if "" == description {
		c.SendView(libview.View{Name: "Todos", Data: map[string]any{
			"todos": state.Todos,
			"error": "todo description cannot be empty",
		}})
		return
	}

	state.Todos = append(state.Todos, lib.Todo{Checked: false, Description: description})

	c.SendView(libview.View{Name: "Todos", Data: map[string]any{
		"todos": state.Todos,
	}})
}
