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
    import Submit from "$frizzante/components/Submit.svelte";
    import Link from '$frizzante/components/Link.svelte'
    import Layout from '$lib/components/Layout.svelte'
    import {getContext} from "svelte";

    /**
     * @typedef Item
     * @property {string} description
     * @property {boolean} checked
     */

    /**
     * @typedef Data
     * @property {Array<Item>} items
     */

    /** @type {Data} */
    const data = getContext("data")
</script>

<Layout title="Todos">
    <div class="items">
        {#each data.items as item, index}
            <div class="item">
                {#if item.checked}
                    <Submit form={{uncheck:index}}>
                        <span class="btn">(x) {item.description}</span>
                    </Submit>
                {:else}
                    <Submit form={{check:index}}>
                        <span class="btn">(&nbsp;&nbsp;) {item.description}</span>
                    </Submit>
                {/if}
            </div>
        {/each}
    </div>
    <br/>
    <div class="menu">
        <Link align="center" view="Welcome">
            <span class="link">&lt; Back</span>
        </Link>
    </div>
</Layout>
