<style>
    a {
        color: inherit;
        text-decoration: inherit;
    }
    a:hover {
        color: inherit;
        text-decoration: inherit;
    }
    a:active {
        color: inherit;
        text-decoration: inherit;
    }
</style>

<script lang="ts">
    import {navigate} from "$lib/scripts/router.ts";
    import type {Snippet} from "svelte";

    type Props = {
        to: string
        children: Snippet
        server: ServerProperties<{}>
        class?: string
        style?: string
    }

    let {
        to,
        children,
        server = $bindable(),
        ...rest
    }: Props = $props()

    async function onclick(e: MouseEvent) {
        e.preventDefault()
        await navigate(server, to)
        return false
    }
</script>

<a href="{server.ids[to]}" {...rest} {onclick}>
    {@render children()}
</a>