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
    import Link from "$frizzante/components/Link.svelte";

    const message = source("/Api/Events", {options: {method: "GET"}}).select("message")
</script>

<Layout title="Welcome">
    <h1>Welcome to Frizzante.</h1>
    <div class="menu">
        <span>{$message}</span><br/>
        <br/>
        <Link page="Todos">
            <span class="link">Show todos</span>
        </Link>
    </div>
</Layout>