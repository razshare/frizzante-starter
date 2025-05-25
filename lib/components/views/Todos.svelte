<style>
    ol {
        padding: 1rem;
        border-radius: 0.3rem;
        background: rgba(0, 0, 0, 0.3);
        list-style-type: none;
        min-width: 400px;
        text-align: start;
    }
</style>

<script lang="ts">
    import Layout from '$lib/components/Layout.svelte'
    import {getContext} from "svelte";
    import type {ServerContext} from "$frizzante/types.ts";
    import Action from "$frizzante/components/Action.svelte";
    import {href} from "$frizzante/scripts/href.ts";

    type Todo = {
        description: string
        checked: boolean
    }

    const server = getContext("server") as ServerContext<{ todos: Todo[] }>
</script>

<Layout title="Todos">
    <ol>
        {#each server.data.todos as item, index}
            <li>
                {#if item.checked}
                    <Action path="/todos" using={{uncheck:index}}>
                        <span class="link">(x) {item.description}</span>
                    </Action>
                {:else}
                    <Action path="/todos" using={{check:index}}>
                        <span class="link">(&nbsp;&nbsp;) {item.description}</span>
                    </Action>
                {/if}
            </li>
        {/each}
    </ol>
    <br/>
    <a class="link" {...href("/")}>&lt; Back</a>
</Layout>
