package handlers

import (
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libsession"
	"github.com/razshare/frizzante/libview"
	"main/lib"
	"strconv"
)

func Uncheck(con *libcon.Connection) {
	state, operator := libsession.Session(con, lib.NewState())
	defer operator.Save(state)

	index := con.ReceiveQuery("index")
	if "" == index {
		return
	}

	id, intError := strconv.ParseInt(index, 10, 64)
	if nil != intError {
		con.SendView(libview.View{Name: "Todos", Data: map[string]any{
			"error": intError.Error(),
		}})
		return
	}

	state.Todos[id].Checked = false

	con.SendView(libview.View{Name: "Todos", Data: map[string]any{
		"todos": state.Todos,
	}})
}
