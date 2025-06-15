package handlers

import (
	"github.com/razshare/frizzante/frz"
	"main/lib"
	"strconv"
)

func Uncheck(c *frz.Connection) {
	state, operator := frz.Session(c, lib.NewState())
	defer operator.Save(state)

	index := c.ReceiveQuery("index")
	if "" == index {
		return
	}

	id, intError := strconv.ParseInt(index, 10, 64)
	if nil != intError {
		c.SendView(frz.View{Name: "Todos", Data: map[string]any{
			"error": intError.Error(),
		}})
		return
	}

	state.Todos[id].Checked = false

	c.SendNavigate("/todos")
}
