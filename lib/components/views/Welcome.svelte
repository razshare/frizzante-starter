<script lang="ts">
    import Layout from '$lib/components/Layout.svelte'
    import {source} from "sveltekit-sse";
    import {href} from "$frizzante/scripts/href.ts";

    const message = source("/api/events", {options: {method: "GET"}}).select("message")
</script>

<Layout title="Welcome">
    <h1>Welcome to Frizzante.</h1>
    <span>{$message}</span><br/>
    <a class="link" {...href("/todos")}>
        Show todos
    </a>
</Layout>