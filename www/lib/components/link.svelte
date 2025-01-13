<style>
    .next {
        color: cadetblue;
        width: 20rem;
        text-align: center;
        cursor: pointer;
        border: 0;
        background: transparent;
    }

    .next:hover {
        text-decoration: underline;
    }
</style>

<script>
    import {getContext} from "svelte";
    const { text, pageId } = $props()
    const page = getContext("page")
</script>
<button class="next" onmouseup={()=>page(pageId)}>{text}</button>