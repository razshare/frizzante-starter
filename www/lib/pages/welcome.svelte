<script>
    import Layout from '$lib/components/layout.svelte'
    import Greeting from '$lib/components/greeting.svelte'
    const { name } = $props()
</script>
<Layout>
    <Greeting {name} />
</Layout>