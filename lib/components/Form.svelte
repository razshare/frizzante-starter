<style>
    form {
        position: relative;
        width: 100%;
        height: 100%;
    }
</style>

<script lang="ts">
    import type {Snippet} from "svelte";
    import {navigate} from "$lib/scripts/router.ts";

    type Props = {
        of: string
        children: Snippet
        server: ServerProperties<{}>
        class?: string
        style?: string
    }

    let {
        of,
        children,
        server = $bindable(),
        ...rest
    }: Props = $props()

    async function onsubmit(e: any) {
        e.preventDefault()
        const form = e.target
        const body = new FormData(form)
        const method = form.method.toUpperCase()
        const headers = {"Accept": "application/json"}
        const response = await fetch(form.action, {method, headers, body})
        if (response.status >= 300) {
            return
        }

        const json = await response.json()

        server.data = {
            ...server.data,
            ...json.data,
        }

        server.ids = {
            ...server.ids,
            ...json.ids
        }

        if (server.id !== json.id) {
            navigate(server, json.id, server.data)
                .then(function done() {
                    server.id = json.id
                })
        }
    }
</script>

<form method="POST" action="{server.ids[of]}" {...rest} {onsubmit}>
    {@render children()}
</form>
