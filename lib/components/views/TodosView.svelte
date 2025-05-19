<style>
    .link {
        color: cadetblue;
        text-decoration: none;
    }

    .link:hover {
        cursor: pointer;
        text-decoration: underline;
    }

    .menu, .item {
        width: 100%;
    }

    .items {
        min-width: 400px;
        padding: 1rem;
        border-radius: 0.3rem;
        background: rgba(0, 0, 0, 0.3);
    }

    .btn {
        color: cadetblue;
        cursor: crosshair;
    }

    .btn:hover {
        text-decoration: underline;
    }
</style>

<script lang="ts">
    import Layout from '$lib/components/Layout.svelte'
    import Action from "$lib/components/Action.svelte";
    import Link from "$lib/components/Link.svelte";
    import Router from "$lib/components/Router.svelte";

    type Item = {
        description: string
        checked: boolean
    }

    type Data = {
        items: Item[]
    }

    type Props = {
        server: ServerProperties<Data>
    }

    let {
        server = $bindable(),
    }: Props = $props()
</script>

<Router bind:server/>
<Layout title="Todos">
    <div class="items">
        {#each server.data.items as item, index}
            <div class="item">
                {#if item.checked}
                    <Action bind:server of="Todos" using={{uncheck:index}}>
                        <span class="btn">(x) {item.description}</span>
                    </Action>
                {:else}
                    <Action bind:server of="Todos" using={{check:index}}>
                        <span class="btn">(&nbsp;&nbsp;) {item.description}</span>
                    </Action>
                {/if}
            </div>
        {/each}
    </div>
    <br/>
    <div class="menu">
        <Link bind:server to="Welcome">
            <span class="link">&lt; Back</span>
        </Link>
    </div>
</Layout>
