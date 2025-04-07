package indexes

import f "github.com/razshare/frizzante"

func Welcome(
	route func(path string, page string),
	_ func(showFunction f.PageFunction),
	_ func(actionFunction f.PageFunction),
) {
	route("/", "welcome")
}
