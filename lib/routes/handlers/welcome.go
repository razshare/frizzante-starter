package handlers

import (
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/send"
	"github.com/razshare/frizzante/view"
)

func Welcome(c *client.Client) {
	send.View(c, view.View{Name: "Welcome"})
}
