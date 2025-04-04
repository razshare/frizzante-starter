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
     * @property {string} Description
     * @property {boolean} Checked
     */

    /**
     * @typedef Data
     * @property {Array<Item>} Items
     */

    /** @type {Data} */
    const data = getContext("Data")
</script>

<Layout title="Todos">
    <div class="items">
        {#each data.Items as item, index}
            <div class="item">
                {#if item.Checked}
                    <Submit form={{Uncheck:index}}>
                        <span class="btn">(x) {item.Description}</span>
                    </Submit>
                {:else}
                    <Submit form={{Check:index}}>
                        <span class="btn">(&nbsp;&nbsp;) {item.Description}</span>
                    </Submit>
                {/if}
            </div>
        {/each}
    </div>
    <br/>
    <div class="menu">
        <Link align="center" pageId="welcome::View">
            <span class="link">&lt; Back</span>
        </Link>
    </div>
</Layout>
