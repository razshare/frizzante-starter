<style>
    .content {
        position: fixed;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        background: #1e1e2e;
        color: #ed6f49;
        display: grid;
        justify-content: center;
        align-content: center;
        font-family: "Noto Sans Gothic",serif;
    }
</style>

<script>
    import Greeting from '$components/greeting.svelte'
    import { data } from "$scripts/data.js";
    const d = data()
</script>

<svelte:head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Welcome</title>
</svelte:head>

<div class="content">
    <Greeting {...d} />
</div>
