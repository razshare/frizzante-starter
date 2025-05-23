package main

import (
	_ "main/lib/api"
	"main/lib/config"
	_ "main/lib/controllers/any"
	_ "main/lib/controllers/expired"
	_ "main/lib/controllers/todos"
)

func main() {
	config.Server.Start()
}
