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

<script lang="ts">
    import Layout from "$lib/components/Layout.svelte";
    import Link from "$lib/components/Link.svelte";
    import Router from "$lib/components/Router.svelte";

    type Props = {
        server: ServerProperties<{}>
    }

    let {
        server = $bindable(),
    }: Props = $props()
</script>

<Router bind:server/>
<Layout title="Expired">
    <h1>Your session has expired!</h1>
    <span class="link">
        <Link bind:server to="Welcome">
            Ok
        </Link>
    </span>
</Layout>