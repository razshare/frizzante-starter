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
    import Router from '$frizzante/components/Router.svelte';
    import Layout from '$lib/components/Layout.svelte'
    import {getContext} from "svelte";
    import type {ServerContext} from "$frizzante/types.ts";
    import Action from "$frizzante/components/Action.svelte";
    import {href} from "$frizzante/scripts/href.ts";

    type Item = {
        description: string
        checked: boolean
    }

    const server = getContext("server") as ServerContext<{ items: Item[] }>
</script>

<Router/>
<Layout title="Todos">
    <div class="items">
        {#each server.data.items as item, index}
            <div class="item">
                {#if item.checked}
                    <Action of="Todos" using={{uncheck:index}}>
                        <span class="btn">(x) {item.description}</span>
                    </Action>
                {:else}
                    <Action of="Todos" using={{check:index}}>
                        <span class="btn">(&nbsp;&nbsp;) {item.description}</span>
                    </Action>
                {/if}
            </div>
        {/each}
    </div>
    <br/>
    <div class="menu">
        <a class="link" {...href("Any")}>&lt; Back</a>
    </div>
</Layout>
