<script lang="ts">
    import {onMount, type Snippet} from "svelte";
    import {route} from "$lib/scripts/router.ts";

    type Props = {
        children: Snippet
        server: ServerProperties<{}>
    }

    let {
        children,
        server = $bindable(),
    }: Props = $props()

    onMount(function () {
        route(server)
    })
</script>

{@render children()}