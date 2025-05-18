<script>
    /**
     * @typedef LinkProperties
     * @property {string} to
     * @property {import("svelte").Snippet} children
     */

    import {navigate} from "$lib/scripts/router.js";

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
        navigate(server, to)
        return false
    }
</script>

<a href="{server.ids[to]}" {onclick}>
    {@render children()}
</a>