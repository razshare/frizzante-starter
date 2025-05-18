<script>
    /**
     * @typedef LinkProperties
     * @property {string} to
     * @property {import("svelte").Snippet} children
     */

    /** @type {ServerProperties<any> & LinkProperties} */
    let {
        server = $bindable(),
        to,
        children,
    } = $props()

    function href() {
        return server.ids[to] ?? ""
    }

    async function onclick(e) {
        e.preventDefault()
        const headers = {"Accept": "application/json"}
        const response = await fetch(`${server.ids[to]}`, {method: "GET", headers})
        if (response.status >= 300) {
            return
        }

        const json = await response.json()

        server.id = json.id
        server.ids = json.ids
        server.data = json.data

        return false
    }
</script>

<a href="{server.ids[to]}" {onclick}>
    {@render children()}
</a>