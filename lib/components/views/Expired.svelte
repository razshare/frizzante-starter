<script lang="ts">
    import Layout from "$lib/components/Layout.svelte";
    import {href} from "$frizzante/scripts/href.ts";
</script>

<Layout title="Session Expired">
    <h1>Your session has expired!</h1>
    <a class="link" {...href("/")}>
        Ok
    </a>
</Layout>