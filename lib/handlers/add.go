package handlers

import (
	"github.com/razshare/frizzante/frz"
	"main/lib"
)

func Add(c *frz.Connection) {
	state, operator := frz.Session(c, lib.NewState())
	defer operator.Save(state)

	description := c.ReceiveQuery("description")

	if "" == description {
		c.SendView(frz.View{Name: "Todos", Data: state.Todos, Error: "todo description cannot be empty"})
		return
	}

	state.Todos = append(state.Todos, lib.Todo{Checked: false, Description: description})

	c.SendView(frz.View{Name: "Todos", Data: state.Todos})
}
