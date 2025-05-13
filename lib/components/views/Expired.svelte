<style>
    .link {
        color: cadetblue;
        text-decoration: none;
    }

    .link:hover {
        cursor: pointer;
        text-decoration: underline;
    }
</style>

<script>
    import {getContext} from "svelte";
    import Layout from "$lib/components/Layout.svelte";
    import Link from "$frizzante/components/Link.svelte";

    /**
     * Server data.
     * @typedef Data
     * @property {any} prop1
     */

    /** @type {Data} */
    const data = getContext("data")
</script>

<Layout title="Expired">
    <h1>Your session has expired!</h1>
    <Link align="center" view="Welcome">
        <span class="link">Ok</span>
    </Link>
</Layout>