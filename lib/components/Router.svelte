<script lang="ts">
    import {onMount} from "svelte";
    import {route} from "$lib/scripts/router.ts";

    type Props = {
        server: ServerProperties<{}>
    }

    let {server = $bindable()}: Props = $props()

    onMount(function () {
        route(server)
    })
</script>