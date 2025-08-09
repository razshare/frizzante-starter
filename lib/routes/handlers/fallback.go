package handlers

import (
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/send"
)

func Fallback(c *client.Client) {
	send.FileOrElse(c, func() { Welcome(c) })
}
