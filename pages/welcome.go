package pages

import frz "github.com/razshare/frizzante"

func Welcome(_ *frz.Server, _ *frz.Request, _ *frz.Response) *frz.SveltePageConfiguration {
	return &frz.SveltePageConfiguration{
		Render: frz.ModeFull,
	}
}
