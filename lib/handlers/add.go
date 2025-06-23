package handlers

import (
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libsession"
	"github.com/razshare/frizzante/libview"
	"main/lib"
)

func Add(con *libcon.Connection) {
	state, operator := libsession.Session(con, lib.NewState())
	defer operator.Save(state)

	description := con.ReceiveQuery("description")

	if "" == description {
		con.SendView(libview.View{Name: "Todos", Data: map[string]any{
			"todos": state.Todos,
			"error": "todo description cannot be empty",
		}})
		return
	}

	state.Todos = append(state.Todos, lib.Todo{Checked: false, Description: description})

	con.SendView(libview.View{Name: "Todos", Data: map[string]any{
		"todos": state.Todos,
	}})
}
