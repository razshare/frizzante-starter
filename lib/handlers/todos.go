package handlers

import (
	"github.com/razshare/frizzante"
	"main/lib"
	"strconv"
)

func GetTodos(c *frizzante.Connection) {
	c.SendView(frizzante.View{Name: "Todos", Data: lib.Todos})
}

type Result[T any] struct {
	Error error
	Value T
}

func PostTodos(c *frizzante.Connection) {
	form := c.ReceiveForm()

	if form.Has("check") {
		id, intError := strconv.ParseInt(form.Get("check"), 10, 64)
		if nil != intError {
			c.SendView(frizzante.View{Name: "Todos", Error: intError})
			return
		}
		lib.Todos[id].Checked = true
	} else if form.Has("uncheck") {
		id, intError := strconv.ParseInt(form.Get("uncheck"), 10, 64)
		if intError != nil {
			c.SendView(frizzante.View{Name: "Todos", Error: intError})
			return
		}
		lib.Todos[id].Checked = false
	}

	c.SendView(frizzante.View{Name: "Todos", Data: lib.Todos})
}
