import { hydrate } from "svelte";
import main from "./main.svelte";
// @ts-ignore
const target = document.getElementById("app");
let props = JSON.parse(atob(target.dataset.props));
hydrate(main, { target, props });
