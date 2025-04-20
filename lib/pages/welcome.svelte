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
    import Link from '$frizzante/components/Link.svelte'
    import Layout from '$lib/components/ui/Layout.svelte'
    import {source} from "sveltekit-sse";
    const message = source("/api/events", {options: {method: "GET"}}).select("message")
</script>

<Layout title="Welcome">
    <h1>Welcome to Frizzante.</h1>
    <div class="menu">
        <span>{$message}</span><br/>
        <br/>
        <Link align="center" page="todos">
            <span class="link">Show todos</span>
        </Link>
    </div>
</Layout>