<script>
    import {getContext} from "svelte";
    import Layout from "$lib/components/Layout.svelte";

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
</Layout>