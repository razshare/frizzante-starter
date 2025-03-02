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
    const path = getContext("path")

    /**
     * @param {Event} e
     */
    function onmouseup(e){
        e.preventDefault()
        page(pageId)
    }
</script>

<a href="{path(pageId)}" class="next" class:start class:center class:end {onmouseup}>{text}</a>