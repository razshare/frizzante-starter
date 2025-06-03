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
    import {href} from "$frizzante/scripts/href.ts";
    import Action from "$frizzante/components/Action.svelte";
    import type {View} from "$frizzante/types.ts";

    type Todo = {
        Checked: boolean
        Description: string
    }

    const view = getContext("view") as View<Todo[]>
</script>

<Layout title="Todos">
    <ol>
        {#each view.data as todo, index(todo.Description+":"+index)}
            <li>
                {#if todo.Checked}
                    <Action path="/todos" using={{uncheck:index}}>
                        <span class="link">(x) {todo.Description}</span>
                    </Action>
                {:else}
                    <Action path="/todos" using={{check:index}}>
                        <span class="link">(&nbsp;&nbsp;) {todo.Description}</span>
                    </Action>
                {/if}
            </li>
        {/each}
    </ol>
    <br/>
    <a class="link" {...href("/")}>&lt; Back</a>
</Layout>
