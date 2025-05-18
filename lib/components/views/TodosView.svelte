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

<script>
    import Layout from '$lib/components/Layout.svelte'
    import Action from "$lib/components/Action.svelte";
    import Link from "$lib/components/Link.svelte";

    /**
     * @typedef Item
     * @property {string} Description
     * @property {boolean} Checked
     */

    /**
     * @typedef Data
     * @property {Array<Item>} Items
     */

    /** @type {ServerProperties<Data>} */
    let {server = $bindable()} = $props()
</script>

<Layout title="Todos">
    <div class="items">
        {#each server.data.Items as item, index}
            <div class="item">
                {#if item.Checked}
                    <Action bind:server of="Todos" using={{uncheck:index}}>
                        <span class="btn">(x) {item.Description}</span>
                    </Action>
                {:else}
                    <Action bind:server of="Todos" using={{check:index}}>
                        <span class="btn">(&nbsp;&nbsp;) {item.Description}</span>
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
