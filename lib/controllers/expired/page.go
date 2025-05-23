package expired

import "main/lib/config"

func init() {
	config.Server.LoadController(nil)
}
