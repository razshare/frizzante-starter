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
</style>

<script>
    const { text, onmouseup } = $props()
</script>
<button class="next" {onmouseup}>{text}</button>