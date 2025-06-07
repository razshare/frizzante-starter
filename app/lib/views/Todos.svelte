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
    <br />
    <a class="link" {...href("/")}>&lt; Back</a>
</Layout>
