<style>
    .next {
        color: cadetblue;
        width: 20rem;
        text-align: start;
        cursor: pointer;
        border: 0;
        background: transparent;
    }

    .next:hover {
        text-decoration: underline;
    }

    .start {
        text-align: start;
    }

    .end {
        text-align: end;
    }
</style>

<script>
    import {getContext} from "svelte";
    const { text, pageId, start, end } = $props()
    const page = getContext("page")
</script>
<button class="next" class:start class:end onmouseup={()=>page(pageId)}>{text}</button>