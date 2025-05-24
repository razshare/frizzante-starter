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

<script lang="ts">
    import Layout from '$lib/components/Layout.svelte'
    // import {source} from "sveltekit-sse";
    import {href} from "$frizzante/scripts/href.ts";

    // const message = source("/api/events", {options: {method: "GET"}}).select("message")
</script>

<Layout title="Welcome">
    <h1>Welcome to Frizzante.</h1>
    <div class="menu">
<!--        <span>{$message}</span><br/>-->
        <br/>
        <a class="link" {...href("todos")}>
            Show todos
        </a>
    </div>
</Layout>