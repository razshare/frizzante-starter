package main

import frz "github.com/razshare/frizzante"

func main() {
	frz.SveltePrepareStart()
	frz.SveltePreparePage("welcome", "./www/lib/pages/welcome.svelte")
	frz.SveltePreparePage("counter", "./www/lib/pages/counter.svelte")
	frz.SveltePrepareEnd()
}
