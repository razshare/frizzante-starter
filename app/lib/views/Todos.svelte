<style>
    ol {
        padding: 1rem;
        border-radius: 0.3rem;
        background: rgba(0, 0, 0, 0.3);
        list-style-type: none;
        min-width: 400px;
        text-align: start;
    }

    input {
        background: transparent;
        border: 0;
        border-bottom: 1px solid cadetblue;
        padding: 0.3rem;
        border-radius: 0;
        color: cadetblue;
    }

    input:focus {
        background-color: rgba(0, 0, 0, 0.1);
        outline: none;
    }

    form {
        display: inline-block;
    }
</style>

<script lang="ts">
    import Layout from "$lib/components/Layout.svelte"
    import { getContext } from "svelte"
    import type { View } from "$lib/utilities/types.ts"
    import { action } from "$lib/utilities/scripts/action.ts"
    import { href } from "$lib/utilities/scripts/href.ts"

    type Todo = {
        Checked: boolean
        Description: string
    }

    const view = getContext("view") as View<Todo[]>
</script>

<Layout title="Todos">
    <ol>
        {#each view.data as todo, index (index)}
            <li>
                <form {...action("/remove")}>
                    <input type="hidden" name="index" value={index} />
                    <button class="link">[Remove]</button>
                </form>
                {#if todo.Checked}
                    <form {...action("/uncheck")}>
                        <input type="hidden" name="index" value={index} />
                        <button class="link">
                            <!---->
                            (x) {todo.Description}
                            <!---->
                        </button>
                    </form>
                {:else}
                    <form {...action("/check")}>
                        <input type="hidden" name="index" value={index} />
                        <button class="link">
                            <!---->
                            (&nbsp;&nbsp;) {todo.Description}
                            <!---->
                        </button>
                    </form>
                {/if}
            </li>
        {/each}
    </ol>
    <form {...action("/add")}>
        <span class="link">Description</span>
        <input type="text" value="" name="description" />
        <button class="link" type="submit">Add +</button>
    </form>

    {#if view.error}
        <br />
        <span class="error">{view.error}</span>
    {/if}

    <br />
    <a class="link" {...href("/")}>&lt; Back</a>
</Layout>
