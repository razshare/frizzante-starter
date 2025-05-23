package expired

import (
	"main/lib/config"
)

func init() {
	config.Server.LoadPageController(nil)
}
