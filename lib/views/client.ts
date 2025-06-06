export const views: Record<string, Promise<unknown>> = {
    "Welcome": import('$lib/components/views/Welcome.svelte'),
    "Todos": import('$lib/components/views/Todos.svelte'),
}