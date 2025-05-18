<style>
    form {
        position: relative;
        width: 100%;
        height: 100%;
    }

    .submit {
        display: none;
    }
</style>

<script>
    import {uuid} from "$lib/scripts/uuid.js";

    const actionId = uuid()

    /**
     * @typedef ActionProperties
     * @property {string} of
     * @property {Record<string,any>} using
     * @property {import("svelte").Snippet} children
     */

    /** @type {ServerProperties<any> & ActionProperties} */
    let {
        server = $bindable(),
        of,
        using = {},
        children,
    } = $props()

    async function onsubmit(e) {
        e.preventDefault()
        /** @type {HTMLFormElement} */
        const form = e.target
        const body = new FormData(form)
        const method = form.method.toUpperCase()
        const headers = {"Accept": "application/json"}
        const response = await fetch(form.action, {method, headers, body})
        if (response.status >= 300) {
            return
        }

        const json = await response.json()

        server.id = json.id
        server.data = json.data
        server.ids = json.ids
    }
</script>

<form method="POST" action="{server.ids[of]}" {onsubmit}>
    {#each Object.keys(using) as key}
        {@const value = using[key]}
        <input type="hidden" name="{key}" value="{value}">
    {/each}

    <input class="submit" type="submit" id="{actionId}"/>

    <label for="{actionId}">
        {@render children()}
    </label>
</form>
