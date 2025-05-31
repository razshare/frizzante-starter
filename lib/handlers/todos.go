package handlers

import (
	frz "github.com/razshare/frizzante"
	"main/lib"
	"strconv"
)

func GetTodos(c *frz.Connection) {
	state, _ := frz.Session(c, lib.NewState())
	c.SendView(frz.View{Name: "Todos", Data: state.Todos})
}

func PostTodos(c *frz.Connection) {
	state, operator := frz.Session(c, lib.NewState())
	defer operator.Save(state)

	form := c.ReceiveForm()

	if form.Has("check") {
		id, intError := strconv.ParseInt(form.Get("check"), 10, 64)
		if nil != intError {
			c.SendView(frz.View{Name: "Todos", Error: intError})
			return
		}
		state.Todos[id].Checked = true
	} else if form.Has("uncheck") {
		id, intError := strconv.ParseInt(form.Get("uncheck"), 10, 64)
		if intError != nil {
			c.SendView(frz.View{Name: "Todos", Error: intError})
			return
		}
		state.Todos[id].Checked = false
	}

	c.SendView(frz.View{Name: "Todos", Data: state.Todos})
}
