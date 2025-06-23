package handlers

import (
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libsession"
	"github.com/razshare/frizzante/libview"
	"main/lib"
	"strconv"
)

func Remove(c *libcon.Connection) {
	state, operator := libsession.Session(c, lib.NewState())
	defer operator.Save(state)

	if 0 == len(state.Todos) {
		// No items found, ignore the request.
		return
	}

	index := c.ReceiveQuery("index")
	if "" == index {
		// No index found, ignore the request.
		return
	}

	id, intError := strconv.ParseInt(index, 10, 64)
	if nil != intError {
		c.SendView(libview.View{Name: "Todos", Data: map[string]any{
			"error": intError.Error(),
		}})
		return
	}

	state.Todos = append(state.Todos[:id], state.Todos[id+1:]...)

	c.SendView(libview.View{Name: "Todos", Data: map[string]any{
		"todos": state.Todos,
	}})
}
