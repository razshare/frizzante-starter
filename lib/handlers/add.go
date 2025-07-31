package handlers

import (
	"github.com/razshare/frizzante/servers"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib"
)

func Add(con *servers.Connection) {
	session := sessions.New(con, lib.InitialState())
	session.Start()
	defer session.Save()

	description := con.ReceiveQuery("description")

	if "" == description {
		con.SendView(views.View{Name: "Todos", Data: map[string]any{
			"todos": session.State.Todos,
			"error": "todo description cannot be empty",
		}})
		return
	}

	session.State.Todos = append(session.State.Todos, lib.Todo{Checked: false, Description: description})

	con.SendNavigate("/todos")
}
