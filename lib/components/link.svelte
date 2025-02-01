<style>
    .next {
        color: cadetblue;
        width: 100%;
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

    .center {
        text-align: center;
    }

    .end {
        text-align: end;
    }
</style>

<script>
    import {getContext} from "svelte";
    const { text, pageId, start = false, center = false, end = false } = $props()
    const page = getContext("page")
</script>

<button class="next" class:start class:center class:end onmouseup={()=>page(pageId)}>{text}</button>