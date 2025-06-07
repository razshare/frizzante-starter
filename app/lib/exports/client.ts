export const views: Record<string, Promise<unknown>> = {
    Welcome: import("$lib/views/Welcome.svelte"),
    Todos: import("$lib/views/Todos.svelte"),
}
