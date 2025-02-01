<style>
    .btn {
        color: cadetblue;
        width: 100%;
        text-align: start;
        cursor: pointer;
        border: 0;
        background: transparent;
    }

    .btn:hover {
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
    const { text, onmouseup, start = false, center = false, end = false } = $props()
</script>
<button class="btn" class:start class:center class:end {onmouseup}>{text}</button>