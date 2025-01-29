<style>
    .menu,.item {
        min-width: 400px;
    }
</style>

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
        <div class="item">
            {#if item.checked}
                <Button start text="(x) {item.description}" onmouseup={()=>toggle(item)}/>
            {:else}
                <Button start text="(&nbsp;&nbsp;) {item.description}" onmouseup={()=>toggle(item)}/>
            {/if}
        </div>
    {/each}
    <br/>
    <div class="menu">
        <Link start text="< Back" pageId="welcome" />
    </div>
</Layout>
