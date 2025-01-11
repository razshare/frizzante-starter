import { render as _render } from "svelte/server";
import {pages} from "./index.js";
export async function render(props) {
  for (const pageName in pages) {
    if(props.page === pageName){
      const component = pages[pageName]
      return _render(component, { props });
    }
  }
}
