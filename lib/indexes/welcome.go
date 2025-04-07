package indexes

import f "github.com/razshare/frizzante"

func Welcome() (
	page string,
	show f.PageFunction,
	action f.PageFunction,
) {
	page = "welcome /"
	return
}
