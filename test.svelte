<script>
    import Page from './page.async.svelte'
    import {setContext} from "svelte";

    let {pageId, paths, path, data} = $props()
    // Do not remove or discard `pageId`, it's being used by app-router.
    let pageIdState = $state(pageId)
    let dataState = $state({...data})
    setContext("data", dataState)
    setContext("page", pageFn)
    setContext("path", pathFn)
    window.addEventListener('popstate', (e) => {
        e.preventDefault()
        if(e.state && e.state.pageId){
            pageFn(e.state.pageId)
        } else {
            for (const pageId in paths) {
                const value = paths[pageId]
                if("/" !== value){
                    continue
                }
                pageFn(pageId)
            }
        }
    });

    function escapeRegExp(string) {
        return string.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')
    }

    function pathFn(pageId) {
        let result = paths[pageId] ?? ''
        if (!paths[pageId]) {
            return ""
        }
        const resolved = {}
        for (let key in path) {
            resolved[key] = false
        }
        for (let key in path) {
            const value = dataState[key]
            const regex = escapeRegExp(`{${key}}`)
            let oldPath = result
            result = result.replaceAll(new RegExp(regex, 'g'), value)
            resolved[key] = oldPath === result
        }

        return result
    }

    function pageFn(pageIdLocal) {
        if (!paths[pageIdLocal]) {
            return
        }

        const pathLocal = pathFn(pageIdLocal)
        history.pushState({pageIdLocal}, '', pathLocal);
        pageIdState = pageIdLocal

        fetch(pathLocal, {headers: {"Accept": "application/json"}}).then(async (response) => {
            const data = await response.json()

            for (const key in dataState) {
                delete dataState[key]
            }

            for (const key in data) {
                dataState[key] = data[key]
            }
        })
    }
</script>

{#if 'Todos' === pageIdState}
    <Page from={import('./../../pages/Todos.svelte')} />
{:else if 'Welcome' === pageIdState}
    <Page from={import('./../../pages/Welcome.svelte')} />
{/if}