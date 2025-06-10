package handlers

import (
	"github.com/razshare/frizzante/frz"
	"main/lib"
	"strconv"
)

func Remove(c *frz.Connection) {
	state, operator := frz.Session(c, lib.NewState())

	if 0 == len(state.Todos) {
		c.SendView(frz.View{Name: "Todos", Data: state.Todos})
		return
	}

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

	if 0 == len(state.Todos) {
		c.SendView(frz.View{Name: "Todos", Data: state.Todos})
		return
	}

	state.Todos = append(state.Todos[:id], state.Todos[id+1:]...)

	c.SendNavigate("/todos")
}
