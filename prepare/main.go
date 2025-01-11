package main

import frz "github.com/razshare/frizzante"

func main() {
	frz.SveltePrepareStart()
	frz.SveltePreparePage("welcome", "./www/lib/pages/welcome.svelte")
	frz.SveltePreparePage("about", "./www/lib/pages/about.svelte")
	frz.SveltePrepareEnd()
}
