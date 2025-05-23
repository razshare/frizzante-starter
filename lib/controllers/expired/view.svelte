<style>
    .link {
        color: cadetblue;
        text-decoration: none;
    }

    .link:hover {
        cursor: pointer;
        text-decoration: underline;
    }
</style>

<script lang="ts">
    import Layout from "$lib/components/Layout.svelte";
    import {href} from "$frizzante/scripts/href.ts";
    import Router from "$frizzante/components/Router.svelte";
</script>

<Router/>
<Layout>
    <h1>Your session has expired!</h1>
    <span class="link">
        <a {...href("any")}>
            Ok
        </a>
    </span>
</Layout>