<script>
    import Layout from '$lib/components/layout.svelte'
    let counter = $state(0)
    function increase() {
        counter++
    }
</script>
<Layout>
    <button onmouseup={increase}>Increase ({counter})</button>
</Layout>