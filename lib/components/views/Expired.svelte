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
    import Layout from "$lib/components/Layout.svelte";
    import Link from "$frizzante/components/Link.svelte";
</script>

<Layout title="Expired">
    <h1>Your session has expired!</h1>
    <Link page="Welcome">
        <span class="link">Ok</span>
    </Link>
</Layout>