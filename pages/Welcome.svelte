<style>
    h1 {
        color: cadetblue;
        text-align: center;
    }
    .menu {
        text-align: center;
    }
    h1,.menu {
        width: 400px;
    }
</style>

<script>
    import Layout from '$lib/components/Layout.svelte'
    import Link from '$lib/components/Link.svelte'
</script>

<svelte:head>
    <title>Welcome</title>
</svelte:head>

<Layout>
    <h1>Welcome to Frizzante.</h1>
    <div class="menu">
        <Link center text="Show todos" pageId="Todos"/>
    </div>
</Layout>