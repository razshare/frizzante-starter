<style>
    .btn {
        color: cadetblue;
        width: 20rem;
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
    const { text, onmouseup, start, end } = $props()
</script>
<button class="btn" class:start class:end {onmouseup}>{text}</button>