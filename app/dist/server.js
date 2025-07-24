var __defProp = Object.defineProperty;
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __publicField = (obj, key, value) => __defNormalProp(obj, typeof key !== "symbol" ? key + "" : key, value);
import { clsx as clsx$1 } from "clsx";
const HYDRATION_START = "[";
const HYDRATION_END = "]";
const ELEMENT_IS_NAMESPACED = 1;
const ELEMENT_PRESERVE_ATTRIBUTE_CASE = 1 << 1;
const ATTR_REGEX = /[&"<]/g;
const CONTENT_REGEX = /[&<]/g;
function escape_html(value, is_attr) {
  const str = String(value ?? "");
  const pattern = is_attr ? ATTR_REGEX : CONTENT_REGEX;
  pattern.lastIndex = 0;
  let escaped = "";
  let last = 0;
  while (pattern.test(str)) {
    const i = pattern.lastIndex - 1;
    const ch = str[i];
    escaped += str.substring(last, i) + (ch === "&" ? "&amp;" : ch === '"' ? "&quot;" : "&lt;");
    last = i + 1;
  }
  return escaped + str.substring(last);
}
const replacements = {
  translate: /* @__PURE__ */ new Map([
    [true, "yes"],
    [false, "no"]
  ])
};
function attr(name, value, is_boolean = false) {
  if (value == null || !value && is_boolean) return "";
  const normalized = name in replacements && replacements[name].get(value) || value;
  const assignment = is_boolean ? "" : `="${escape_html(normalized, true)}"`;
  return ` ${name}${assignment}`;
}
function clsx(value) {
  if (typeof value === "object") {
    return clsx$1(value);
  } else {
    return value ?? "";
  }
}
function to_class(value, hash, directives) {
  var classname = value == null ? "" : "" + value;
  if (hash) {
    classname = classname ? classname + " " + hash : hash;
  }
  return classname === "" ? null : classname;
}
const DEV = false;
const STALE_REACTION = new class StaleReactionError extends Error {
  constructor() {
    super(...arguments);
    __publicField(this, "name", "StaleReactionError");
    __publicField(this, "message", "The reaction that called `getAbortSignal()` was re-run or destroyed");
  }
}();
function lifecycle_outside_component(name) {
  {
    throw new Error(`https://svelte.dev/e/lifecycle_outside_component`);
  }
}
const DOM_BOOLEAN_ATTRIBUTES = [
  "allowfullscreen",
  "async",
  "autofocus",
  "autoplay",
  "checked",
  "controls",
  "default",
  "disabled",
  "formnovalidate",
  "hidden",
  "indeterminate",
  "inert",
  "ismap",
  "loop",
  "multiple",
  "muted",
  "nomodule",
  "novalidate",
  "open",
  "playsinline",
  "readonly",
  "required",
  "reversed",
  "seamless",
  "selected",
  "webkitdirectory",
  "defer",
  "disablepictureinpicture",
  "disableremoteplayback"
];
function is_boolean_attribute(name) {
  return DOM_BOOLEAN_ATTRIBUTES.includes(name);
}
var current_component = null;
function getContext(key) {
  const context_map = get_or_init_context_map();
  const result = (
    /** @type {T} */
    context_map.get(key)
  );
  return result;
}
function setContext(key, context) {
  get_or_init_context_map().set(key, context);
  return context;
}
function get_or_init_context_map(name) {
  if (current_component === null) {
    lifecycle_outside_component();
  }
  return current_component.c ?? (current_component.c = new Map(get_parent_context(current_component) || void 0));
}
function push(fn) {
  current_component = { p: current_component, c: null, d: null };
}
function pop() {
  var component = (
    /** @type {Component} */
    current_component
  );
  var ondestroy = component.d;
  if (ondestroy) {
    on_destroy.push(...ondestroy);
  }
  current_component = component.p;
}
function get_parent_context(component_context) {
  let parent = component_context.p;
  while (parent !== null) {
    const context_map = parent.c;
    if (context_map !== null) {
      return context_map;
    }
    parent = parent.p;
  }
  return null;
}
const BLOCK_OPEN = `<!--${HYDRATION_START}-->`;
const BLOCK_CLOSE = `<!--${HYDRATION_END}-->`;
class HeadPayload {
  constructor(css = /* @__PURE__ */ new Set(), out = [], title = "", uid = () => "") {
    /** @type {Set<{ hash: string; code: string }>} */
    __publicField(this, "css", /* @__PURE__ */ new Set());
    /** @type {string[]} */
    __publicField(this, "out", []);
    __publicField(this, "uid", () => "");
    __publicField(this, "title", "");
    this.css = css;
    this.out = out;
    this.title = title;
    this.uid = uid;
  }
}
class Payload {
  constructor(id_prefix = "") {
    /** @type {Set<{ hash: string; code: string }>} */
    __publicField(this, "css", /* @__PURE__ */ new Set());
    /** @type {string[]} */
    __publicField(this, "out", []);
    __publicField(this, "uid", () => "");
    __publicField(this, "select_value");
    __publicField(this, "head", new HeadPayload());
    this.uid = props_id_generator(id_prefix);
    this.head.uid = this.uid;
  }
}
function props_id_generator(prefix) {
  let uid = 1;
  return () => `${prefix}s${uid++}`;
}
function reset_elements() {
  return () => {
  };
}
let controller = null;
function abort() {
  controller == null ? void 0 : controller.abort(STALE_REACTION);
  controller = null;
}
const INVALID_ATTR_NAME_CHAR_REGEX = /[\s'">/=\u{FDD0}-\u{FDEF}\u{FFFE}\u{FFFF}\u{1FFFE}\u{1FFFF}\u{2FFFE}\u{2FFFF}\u{3FFFE}\u{3FFFF}\u{4FFFE}\u{4FFFF}\u{5FFFE}\u{5FFFF}\u{6FFFE}\u{6FFFF}\u{7FFFE}\u{7FFFF}\u{8FFFE}\u{8FFFF}\u{9FFFE}\u{9FFFF}\u{AFFFE}\u{AFFFF}\u{BFFFE}\u{BFFFF}\u{CFFFE}\u{CFFFF}\u{DFFFE}\u{DFFFF}\u{EFFFE}\u{EFFFF}\u{FFFFE}\u{FFFFF}\u{10FFFE}\u{10FFFF}]/u;
let on_destroy = [];
function render$1(component, options = {}) {
  try {
    const payload = new Payload(options.idPrefix ? options.idPrefix + "-" : "");
    const prev_on_destroy = on_destroy;
    on_destroy = [];
    payload.out.push(BLOCK_OPEN);
    let reset_reset_element;
    if (DEV) ;
    if (options.context) {
      push();
      current_component.c = options.context;
    }
    component(payload, options.props ?? {}, {}, {});
    if (options.context) {
      pop();
    }
    if (reset_reset_element) {
      reset_reset_element();
    }
    payload.out.push(BLOCK_CLOSE);
    for (const cleanup of on_destroy) cleanup();
    on_destroy = prev_on_destroy;
    let head2 = payload.head.out.join("") + payload.head.title;
    for (const { hash, code } of payload.css) {
      head2 += `<style id="${hash}">${code}</style>`;
    }
    const body = payload.out.join("");
    return {
      head: head2,
      html: body,
      body
    };
  } finally {
    abort();
  }
}
function head(payload, fn) {
  const head_payload = payload.head;
  head_payload.out.push(BLOCK_OPEN);
  fn(head_payload);
  head_payload.out.push(BLOCK_CLOSE);
}
function spread_attributes(attrs, css_hash, classes, styles, flags = 0) {
  if (attrs.class) {
    attrs.class = clsx(attrs.class);
  }
  if (css_hash || classes) {
    attrs.class = to_class(attrs.class, css_hash);
  }
  let attr_str = "";
  let name;
  const is_html = (flags & ELEMENT_IS_NAMESPACED) === 0;
  const lowercase = (flags & ELEMENT_PRESERVE_ATTRIBUTE_CASE) === 0;
  for (name in attrs) {
    if (typeof attrs[name] === "function") continue;
    if (name[0] === "$" && name[1] === "$") continue;
    if (INVALID_ATTR_NAME_CHAR_REGEX.test(name)) continue;
    var value = attrs[name];
    if (lowercase) {
      name = name.toLowerCase();
    }
    attr_str += attr(name, value, is_html && is_boolean_attribute(name));
  }
  return attr_str;
}
function spread_props(props) {
  const merged_props = {};
  let key;
  for (let i = 0; i < props.length; i++) {
    const obj = props[i];
    for (key in obj) {
      const desc = Object.getOwnPropertyDescriptor(obj, key);
      if (desc) {
        Object.defineProperty(merged_props, key, desc);
      } else {
        merged_props[key] = obj[key];
      }
    }
  }
  return merged_props;
}
function ensure_array_like(array_like_or_iterator) {
  if (array_like_or_iterator) {
    return array_like_or_iterator.length !== void 0 ? array_like_or_iterator : Array.from(array_like_or_iterator);
  }
  return [];
}
const $$css$2 = {
  hash: "svelte-dz4ykl",
  code: '.content.svelte-dz4ykl {position:fixed;left:0;right:0;top:0;bottom:0;background:#1e1e2e;color:#ed6f49;display:grid;justify-content:center;align-content:center;font-family:"Noto Sans Gothic", serif;text-align:center;}'
};
function Layout($$payload, $$props) {
  $$payload.css.add($$css$2);
  push();
  const view = getContext("view");
  let { title = view.name, children } = $$props;
  head($$payload, ($$payload2) => {
    $$payload2.title = `<title>${escape_html(title)}</title>`;
    $$payload2.out.push(`<meta charset="UTF-8"/> <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0"/>`);
  });
  $$payload.out.push(`<div class="content svelte-dz4ykl">`);
  children($$payload);
  $$payload.out.push(`<!----></div>`);
  pop();
}
const IS_BROWSER = typeof document !== "undefined";
function uuid(short = false) {
  let dt = (/* @__PURE__ */ new Date()).getTime();
  let blueprint = "xxxxxxxx-xxxx-yxxx-yxxx-xxxxxxxxxxxx";
  if (short) {
    blueprint = "xyxxyxyx";
  }
  return blueprint.replace(/[xy]/g, function check(c) {
    const r = (dt + Math.random() * 16) % 16 | 0;
    dt = Math.floor(dt / 16);
    return (c === "x" ? r : r & 3 | 8).toString(16);
  });
}
let nextPosition = 0;
const record = {};
function find(id) {
  return record[id] ?? false;
}
function swap(view) {
  let swapMethod = "GET";
  let swapPath = location.pathname;
  let swapBody;
  const swapPosition = nextPosition++;
  return {
    method() {
      return swapMethod;
    },
    path() {
      return swapPath;
    },
    body() {
      return swapBody;
    },
    position() {
      return swapPosition;
    },
    withMethod(method) {
      swapMethod = method.toUpperCase();
      return this;
    },
    withPath(path) {
      swapPath = path;
      return this;
    },
    withBody(body) {
      swapBody = body;
      return this;
    },
    async play(update) {
      const payload = {
        method: swapMethod,
        headers: { Accept: "application/json" }
      };
      let query = "";
      if ("GET" === swapMethod) {
        if (swapBody && typeof swapBody === "object") {
          const params = new URLSearchParams();
          swapBody.forEach(function each(value, key) {
            params.append(key, `${value}`);
          });
          query = `${params.toString()}`;
          if (swapPath.includes("?")) {
            query = "&" + query;
          } else {
            query = "?" + query;
          }
        }
      } else {
        payload.body = swapBody;
      }
      const response = await fetch(`${swapPath}${query}`, payload);
      const text = await response.text();
      if ("" === text) {
        return;
      }
      const json = JSON.parse(text);
      view.data = json.data;
      view.name = json.name;
      view.renderMode = json.renderMode;
      if (update) {
        const id = uuid();
        record[id] = this;
        window.history.pushState(id, "", response.url);
      }
    }
  };
}
function position() {
  return nextPosition;
}
function teleport(position2) {
  nextPosition = position2;
}
const swaps = {
  swap,
  find,
  position,
  teleport
};
let started = false;
function route(view) {
  if (!IS_BROWSER || started) {
    return;
  }
  const listener = async function pop2(e) {
    e.preventDefault();
    const id = e.state ?? "";
    const current = swaps.find(id);
    if (!current) {
      await swaps.swap(view).withPath("/").play(false);
      return;
    }
    if (current.position() + 1 != swaps.position()) {
      swaps.teleport(current.position() + 1);
      await current.play(false);
    } else {
      await current.play(true);
    }
  };
  window.addEventListener("popstate", listener);
  started = true;
}
function href(path = "") {
  const view = getContext("view");
  route(view);
  return {
    href: path,
    async onclick(event) {
      event.preventDefault();
      await swaps.swap(view).withPath(path).play(true);
      return false;
    }
  };
}
function Link($$payload, $$props) {
  push();
  let { href: path, children, class: cls, style } = $$props;
  let pending = false;
  let error = false;
  let options = function run() {
    const out = href(path);
    return {
      href: out.href,
      onclick(event) {
        pending = true;
        out.onclick(event).then(function run2() {
          pending = false;
        }).catch(function run2(errorLocal) {
          error = errorLocal;
        });
      }
    };
  }();
  $$payload.out.push(`<a${spread_attributes({ ...options, class: clsx(cls), style }, null)}>`);
  children($$payload, { pending, error });
  $$payload.out.push(`<!----></a>`);
  pop();
}
function Welcome($$payload) {
  Layout($$payload, {
    title: "Welcome",
    children: ($$payload2) => {
      $$payload2.out.push(`<h1>Welcome to Frizzante.</h1> `);
      Link($$payload2, {
        class: "link",
        href: "/todos",
        children: ($$payload3) => {
          $$payload3.out.push(`<span>Show todos</span>`);
        }
      });
      $$payload2.out.push(`<!---->`);
    }
  });
}
function action(path = "") {
  const view = getContext("view");
  route(view);
  return {
    action: path,
    async onsubmit(event) {
      event.preventDefault();
      const form = event.target;
      const body = new FormData(form);
      const target = event.target;
      await swaps.swap(view).withMethod(target.method).withPath(path).withBody(body).play(true).then(function done() {
        form.reset();
      });
    }
  };
}
const $$css$1 = {
  hash: "svelte-bzpzdn",
  code: "form.svelte-bzpzdn {display:inline-block;}"
};
function Form($$payload, $$props) {
  $$payload.css.add($$css$1);
  push();
  let {
    method = "GET",
    action: actionPath,
    children,
    class: cls,
    style
  } = $$props;
  let pending = false;
  let error = false;
  let options = function run() {
    const out = action(actionPath);
    return {
      action: out.action,
      onsubmit(event) {
        pending = true;
        out.onsubmit(event).then(function run2() {
          pending = false;
        }).catch(function run2(errorLocal) {
          error = errorLocal;
        });
      }
    };
  }();
  $$payload.out.push(`<form${spread_attributes({ method, ...options, class: clsx(cls), style }, "svelte-bzpzdn")}>`);
  children($$payload, { pending, error });
  $$payload.out.push(`<!----></form>`);
  pop();
}
const $$css = {
  hash: "svelte-1gqr49j",
  code: "ol.svelte-1gqr49j {padding:1rem;border-radius:0.3rem;background:rgba(0, 0, 0, 0.3);list-style-type:none;min-width:400px;text-align:start;}input.svelte-1gqr49j {background:transparent;border:0;border-bottom:1px solid cadetblue;padding:0.3rem;border-radius:0;color:cadetblue;}input.svelte-1gqr49j:focus {background-color:rgba(0, 0, 0, 0.1);outline:none;}"
};
function Todos($$payload, $$props) {
  $$payload.css.add($$css);
  push();
  let { todos, error } = $$props;
  Layout($$payload, {
    title: "Todos",
    children: ($$payload2) => {
      const each_array = ensure_array_like(todos);
      $$payload2.out.push(`<ol class="svelte-1gqr49j"><!--[-->`);
      for (let index = 0, $$length = each_array.length; index < $$length; index++) {
        let todo = each_array[index];
        $$payload2.out.push(`<li>`);
        Form($$payload2, {
          action: "/remove",
          children: ($$payload3) => {
            $$payload3.out.push(`<input type="hidden" name="index"${attr("value", index)} class="svelte-1gqr49j"/> <button class="link">[Remove]</button>`);
          }
        });
        $$payload2.out.push(`<!----> `);
        if (todo.Checked) {
          $$payload2.out.push("<!--[-->");
          Form($$payload2, {
            action: "/uncheck",
            children: ($$payload3) => {
              $$payload3.out.push(`<input type="hidden" name="index"${attr("value", index)} class="svelte-1gqr49j"/> <button class="link">(x) ${escape_html(todo.Description)}</button>`);
            }
          });
        } else {
          $$payload2.out.push("<!--[!-->");
          Form($$payload2, {
            action: "/check",
            children: ($$payload3) => {
              $$payload3.out.push(`<input type="hidden" name="index"${attr("value", index)} class="svelte-1gqr49j"/> <button class="link">(  ) ${escape_html(todo.Description)}</button>`);
            }
          });
        }
        $$payload2.out.push(`<!--]--></li>`);
      }
      $$payload2.out.push(`<!--]--></ol> `);
      Form($$payload2, {
        action: "/add",
        children: ($$payload3) => {
          $$payload3.out.push(`<span class="link">Description</span> <input type="text" value="" name="description" class="svelte-1gqr49j"/> <button class="link" type="submit">Add +</button>`);
        }
      });
      $$payload2.out.push(`<!----> `);
      if (error) {
        $$payload2.out.push("<!--[-->");
        $$payload2.out.push(`<br/> <span class="error">${escape_html(error)}</span>`);
      } else {
        $$payload2.out.push("<!--[!-->");
      }
      $$payload2.out.push(`<!--]--> <br/> <a${spread_attributes({ class: "link", ...href("/") }, null)}>&lt; Back</a>`);
    }
  });
  pop();
}
const views = {
  Welcome,
  Todos
};
function ServerView($$payload, $$props) {
  push();
  const components = views;
  let { name, data, renderMode } = $$props;
  const view = { name, data, renderMode };
  setContext("view", view);
  const each_array = ensure_array_like(Object.keys(components));
  $$payload.out.push(`<!--[-->`);
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let key = each_array[$$index];
    const Component = components[key];
    if (key === name) {
      $$payload.out.push("<!--[-->");
      $$payload.out.push(`<!---->`);
      Component($$payload, spread_props([view.data]));
      $$payload.out.push(`<!---->`);
    } else {
      $$payload.out.push("<!--[!-->");
    }
    $$payload.out.push(`<!--]-->`);
  }
  $$payload.out.push(`<!--]-->`);
  pop();
}
async function render(props) {
  return render$1(ServerView, { props });
}
export {
  render
};
