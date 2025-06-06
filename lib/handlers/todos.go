package handlers

import (
	"github.com/razshare/frizzante/frz"
	"main/lib"
	"strconv"
)

func Todos(c *frz.Connection) {
	state, operator := frz.Session(c, lib.NewState())
	defer operator.Save(state)

	index := c.ReceiveQuery("index")
	if "" != index {
		id, intError := strconv.ParseInt(index, 10, 64)
		if nil != intError {
			c.SendView(frz.View{Name: "Todos", Error: intError.Error()})
			return
		}
		action := c.ReceiveQuery("action")
		if "check" == action {
			state.Todos[id].Checked = true
		} else if "uncheck" == action {
			state.Todos[id].Checked = false
		}
	}

	c.SendView(frz.View{Name: "Todos", Data: state.Todos})
}
