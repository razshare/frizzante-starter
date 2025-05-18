<style>
    .link {
        color: cadetblue;
        text-decoration: none;
    }

    .link:hover {
        cursor: pointer;
        text-decoration: underline;
    }

    h1 {
        color: cadetblue;
        text-align: center;
        width: 400px;
    }

    .menu {
        text-align: center;
        width: 400px;
    }
</style>

<script>
    import Layout from '$lib/components/Layout.svelte'
    import {source} from "sveltekit-sse";
    import Link from "$lib/components/Link.svelte";

    const message = source("/api/events", {options: {method: "GET"}}).select("message")

    /** @type {ServerProperties<any>} */
    let {server = $bindable()} = $props()
</script>

<Layout title="Welcome">
    <h1>Welcome to Frizzante.</h1>
    <div class="menu">
        <span>{$message}</span><br/>
        <br/>
        <Link bind:server to="Todos">
            <span class="link">Show todos</span>
        </Link>
    </div>
</Layout>