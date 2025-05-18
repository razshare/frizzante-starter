<style global>
    .content {
        position: fixed;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        background: #1e1e2e;
        color: #ed6f49;
        display: grid;
        justify-content: center;
        align-content: center;
        font-family: "Noto Sans Gothic", serif;
    }
</style>

<script>
    /**
     * @typedef LayoutProperties
     * @property {string} title
     * @property {import("svelte").Snippet} children
     */

    import Router from "$lib/components/Router.svelte";

    /** @type {ServerProperties<any> & LayoutProperties} */
    let { server = $bindable(), title, children} = $props()
</script>

<svelte:head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <title>{title}</title>
</svelte:head>

<Router bind:server>
    <div class="content">
        {@render children()}
    </div>
</Router>