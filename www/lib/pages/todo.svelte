<script>
    import Layout from '$lib/components/layout.svelte'
    import Button from "$lib/components/button.svelte";
    import Link from "$lib/components/link.svelte";
    import {getContext} from "svelte";
    const data = getContext("data")
    function toggle(item){
        item.checked = !item.checked

        fetch("/check", {
            method: 'POST',
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data.items)
        })
    }
</script>

<svelte:head>
    <title>Todo</title>
</svelte:head>

<Layout>
    {#each data.items as item}
        <span>
            {#if item.checked}
                <Button text="(x) {item.description}" onmouseup={()=>toggle(item)}/>
            {:else}
                <Button text="( ) {item.description}" onmouseup={()=>toggle(item)}/>
            {/if}
        </span>
    {/each}
    <br/>
    <Link text="Back" pageId="welcome" />
</Layout>
