import { hydrate } from "svelte";
import {pages} from "./index.js";
const p = props()
for (const pageName in pages) {
    if(p.page === pageName){
        const component = pages[pageName]
        // @ts-ignore
        hydrate(component, { target: target(), props: p});
        break
    }
}

