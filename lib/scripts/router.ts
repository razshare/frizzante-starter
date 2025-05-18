type Modifier = "push" | "back" | "forward"

let counter = 0;

async function swap(server: ServerProperties<any>, id: string, modifier: Modifier, data: any = false): Promise<void> {
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
    server.ids = {
        ...server.ids,
        ...json.ids,
    }
    server.id = json.id;
}


export function navigate(server: ServerProperties<any>, id: string, data: any = false): Promise<void> {
    return swap(server, id, "push", data);
}

export function route(server: ServerProperties<any>): void {
    const listener = async function pop(e: PopStateEvent) {
        e.preventDefault();
        const idLocal = e.state?.id ?? "";
        const counterLocal = e.state?.counter ?? 0;
        if (counterLocal < counter) {
            counter = counterLocal;
            await swap(server, idLocal, "back");
        } else if (counterLocal > counter) {
            counter = counterLocal;
            await swap(server, idLocal, "forward");
        } else {
            await swap(server, idLocal, "push");
        }
    }
    window.addEventListener("popstate", listener);
}
