<script>
    import Layout from '$lib/components/layout.svelte'
    import Greeting from '$lib/components/greeting.svelte'
    const { name = "world" } = $props()
</script>

<svelte:head>
    <title>Welcome</title>
</svelte:head>

<Layout>
    <Greeting {name} />
    <a href="/about">Go to about</a>
</Layout>