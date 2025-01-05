import { hydrate } from "svelte";
import main from "./main.svelte";
// @ts-ignore
hydrate(main, { target: target(), props: props() });
