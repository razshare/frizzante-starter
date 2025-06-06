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
    import {href} from "$frz/scripts/href.ts";
    import type {View} from "$frz/types.ts";
    import {action} from "$frz/scripts/action.ts";

    type Todo = {
        Checked: boolean
        Description: string
    }

    const view = getContext("view") as View<Todo[]>
</script>

<Layout title="Todos">
    <ol>
        {#each view.data as todo, index(todo.Description + ":" + index)}
            <li>
                <form method="GET" {...action("/todos?")}>
                    <input type="hidden" name="index" value="{index}"/>
                    {#if todo.Checked}
                        <input type="hidden" name="action" value="uncheck"/>
                        <button class="link">(x) {todo.Description}</button>
                    {:else}
                        <input type="hidden" name="action" value="check"/>
                        <button class="link">(&nbsp;&nbsp;) {todo.Description}</button>
                    {/if}
                </form>
            </li>
        {/each}
    </ol>
    <br/>
    <a class="link" {...href("/")}>&lt; Back</a>
</Layout>
