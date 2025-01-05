import { render as _render } from "svelte/server";
import main from "./main.svelte";
export async function render(props) {
  return _render(main, { props });
}
