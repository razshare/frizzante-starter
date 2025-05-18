let counter = 0;

/**
 *
 * @param {Server<any>} server
 * @param {string} id
 * @param {"back"|"forward"|"push"} modifier
 * @param {false|Record<string,any>} [data]
 * @returns {Promise<void>}
 */
async function swap(server, id, modifier, data = false) {
    if (!server.ids[id]) {
        return;
    }

    const path = server.ids[id];

    if ("push" === modifier) {
        window.history.pushState({id, counter: ++counter}, "", path);
    }

    if (false !== data) {
        server.id = id;
        return;
    }

    const response = await fetch(path, {
        headers: {Accept: "application/json"},
    });
    const json = await response.json();
    server.data = {
        ...server.data,
        ...json.data,
    }
    server.id = json.id;
}

/**
 *
 * @param {Server<any>} server
 * @param {string} id
 * @param {false|Record<string,any>} [data]
 * @returns {Promise<void>}
 */
export function navigate(server, id, data = false) {
    return swap(server, id, "push", data);
}

/**
 *
 * @param {Server<any>} server
 */
export function route(server) {
    const listener = async function pop(e) {
        e.preventDefault();
        const viewLocal = e.state?.id ?? "";
        const counterLocal = e.state?.counter ?? 0;
        if (counterLocal < counter) {
            swap(server, viewLocal, "back");
            counter = counterLocal;
        } else if (counterLocal > counter) {
            swap(server, viewLocal, "forward");
            counter = counterLocal;
        } else {
            swap(server, viewLocal, "push");
        }
    }
    window.addEventListener("popstate", listener);
}
