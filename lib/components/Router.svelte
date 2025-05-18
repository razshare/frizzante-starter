<script>
    /**
     * @typedef Props
     * @property {import("svelte").Snippet} children
     */

    import {onMount} from "svelte";
    import {route} from "$lib/scripts/router.js";

    /** @type {ServerProperties<any> & Props} */
    let {server = $bindable(), children} = $props()

    onMount(function () {
        route(server)
    })
</script>

{@render children()}