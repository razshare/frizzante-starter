package handlers

import (
	"github.com/razshare/frizzante/frz"
	"main/lib"
	"strconv"
)

func Check(c *frz.Connection) {
	state, operator := frz.Session(c, lib.NewState())
	defer operator.Save(state)

	index := c.ReceiveQuery("index")
	if "" == index {
		return
	}

	id, intError := strconv.ParseInt(index, 10, 64)
	if nil != intError {
		c.SendView(frz.View{Name: "Todos", Error: intError.Error()})
		return
	}

	state.Todos[id].Checked = true

	c.SendView(frz.View{Name: "Todos", Data: state.Todos})
}
