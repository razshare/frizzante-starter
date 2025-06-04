export const views: Record<string, Promise<unknown>> = {
    "Welcome": import('./Welcome.svelte'),
    "Todos": import('./Todos.svelte'),
}