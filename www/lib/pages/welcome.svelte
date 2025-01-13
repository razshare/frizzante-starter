<style>
    h1 {
        color: cadetblue;
    }
</style>

<script>
    import Layout from '$lib/components/layout.svelte'
    import Link from '$lib/components/link.svelte'
</script>

<svelte:head>
    <title>Welcome</title>
</svelte:head>

<Layout>
    <h1>Welcome to Frizzante.</h1>
    <Link text="Next" pageId="todo"/>
</Layout>